pipeline {
    environment {
        HOME = '.'
    }
    agent {
        docker {
            image 'node:20-alpine'
//             args  '-v /tmp:/tmp'
//             args '-v ./:/app'
//             args '-v /var/lib/jenkins/docker_volumes/yudoleplatform/chatserver/front/node_modules:${WORKSPACE}/front/node_modules'
        }
    }

    stages {
//         stage('Git') {
//             steps {
//                 git branch: 'develop',
//                     url: 'https://github.com/WebmiRU/jenkins-test-1.git'
//             }
//         }

        stage('Build front') {
            steps {
                sh 'printenv'
                dir('frontend') {
                    sh 'rm -f package-lock.json'
                    sh 'rm -fr dist'
                    sh 'npm install'
                    sh 'npm run build'
                    sh 'ls -al'
                    sh 'ls -al dist'
                }
            }
        }
    }
//     post {
//         always {
//             archiveArtifacts artifacts: 'theme1/dist/**', fingerprint: true
//         }
//     }
}
