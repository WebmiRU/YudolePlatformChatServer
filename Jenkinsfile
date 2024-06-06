pipeline {
    environment {
        HOME = '.'
    }
    agent {
        docker {
            image 'node:20-alpine'
//             args  '-v /tmp:/tmp'
//             args '-v ./:/app'
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
                dir('frontend') {
                    sh 'rm -f package-lock.json'
                    sh 'rm -fr dist'
                    sh 'npm install'
                    sh 'npm run build'
                    sh 'ls -al'
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
