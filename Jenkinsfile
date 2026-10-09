pipeline {
  agent any                                   // 在 Jenkins 本机上运行
  triggers { pollSCM('H/2 * * * *') }         // 每 2 分钟检查一次仓库，有新提交才构建
  options {
    disableConcurrentBuilds()                 // 同一时间只跑一个构建
    buildDiscarder(logRotator(numToKeepStr: '10')) //只保留最近 10 次构建记录，更早的自动删掉。
  }
  environment {
    IMAGE       = '192.168.56.13:5000/labapp'
    GITOPS_REPO = 'github.com/roccowork/labapp-gitops.git'
  }
  stages {
    stage('Build & Test') {
      steps {
        script {
          // tag = 构建号-提交号前 7 位，比如 7-a1b2c3d：能对应到具体哪次构建、哪次提交
          env.TAG = "${env.BUILD_NUMBER}-${env.GIT_COMMIT.substring(0, 7)}"
        }
        sh 'docker build --build-arg APP_VERSION=$TAG -t $IMAGE:$TAG .'
      }
    }
    stage('Push') {
      steps {
        sh 'docker push $IMAGE:$TAG'
      }
    }
    stage('Update GitOps repo') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'github-token',
                                          usernameVariable: 'GH_USER', passwordVariable: 'GH_TOKEN')]) {
          sh '''
            rm -rf gitops
            git clone https://$GH_USER:$GH_TOKEN@$GITOPS_REPO gitops
            cd gitops
            sed -i "s#image: 192.168.56.13:5000/labapp:[^ ]*#image: $IMAGE:$TAG#" labapp/deployment.yaml
            git diff
            git config user.name  jenkins
            git config user.email jenkins@lab.local
            git commit -am "Deploy labapp $TAG"
            git push
          '''
        }
      }
    }
  }
  post {
    always { sh 'docker image prune -f' }     // 清理构建产生的无用镜像层，省磁盘
  }
}