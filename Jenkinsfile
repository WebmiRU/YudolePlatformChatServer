pipeline {
    agent none

    environment {
        HOME = '.'
    }

    stages {
//         stage('Git') {
//             steps {
//                 git branch: 'develop',
//                     url: 'https://github.com/WebmiRU/jenkins-test-1.git'
//             }
//         }

//         stage('Build front') {
//             agent {
//                 docker {
//                     image 'node:20-alpine'
//                     reuseNode true
//         //             args  '-v /tmp:/tmp'
//         //             args '-v ./:/app'
//         //             args '-v /var/lib/jenkins/docker_volumes/yudoleplatform/chatserver/front/node_modules:${WORKSPACE}/front/node_modules'
//                 }
//             }
//
//             steps {
// //                 sh 'printenv'
//                 dir('frontend') {
//                     sh 'rm -f package-lock.json'
// //                     sh 'rm -fr dist'
// //                     sh 'npm install'
//                     sh 'npm run build || npm install && npm run build'
//                     sh 'ls -al'
//                     sh 'ls -al dist/assets'
//                 }
//             }
//         }
//
//         stage('Build themes') {
//             agent {
//                 docker {
//                     image 'node:20-alpine'
//                     reuseNode true
//                 }
//             }
//
//             steps {
// //                 sh 'printenv'
//                 dir('themes') {
//                     sh '''
//                         for theme in ./*
//                         do
//                             if [ -d ${theme} ]; then
//                                 cd ${theme}
//                                 rm -f package-lock.json
//                                 npm run build || npm install && npm run build
//                             fi
//                         done
//                     '''
//                 }
//             }
//         }

        stage('Build Win64/exe') {
            agent {
                docker {
                    image 'golang:1.22-alpine'
                    reuseNode true
                    args '-v :${WORKSPACE}/.cache/mods:${WORKSPACE}/.cache/mods'
                }
            }

            steps {
                environment {
                    GOOS = windows
                    GOARCH = amd64
                    GOMODCACHE = {WORKSPACE}/.cache/mods
                }
                sh 'go build -buildvcs=false -o chatserver.exe'
                sh 'ls -al'
            }
        }
    }
//     post {
//         always {
//             archiveArtifacts artifacts: 'theme1/dist/**', fingerprint: true
//         }
//     }
}
