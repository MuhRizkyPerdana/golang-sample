pipeline {
    agent any
    environment {
        TRAINER_PRIVATE_KEY = credentials('private-key-trainer')
    }
    stages {
        stage('Build') {
            agent {
                docker { image 'golang:alpine3.16' }
            }
            steps {
                sh 'HOME=${WORKSPACE} GOOS=darwin GOARCH=arm64 go build -o golang-sample-macos-arm64 main.go'
            }
        }
        stage('Deploy to server') {
            agent {
                dockerfile {
                    filename 'Dockerfile.build'
                }
            }
            steps {
                sh 'ls'   
                sh 'ssh -i ${TRAINER_PRIVATE_KEY} trainer@10.184.15.233 whoami'
            }
        }
    }
    post {
        always {
            archiveArtifacts artifacts: 'golang-sample-macos-arm64', fingerprint: true
        }
    }
}
