pipeline {
    agent none
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
                docker { image 'alpine:latest' }
            }
            steps {
                sh 'echo hai'   
            }
        }
    }
    post {
        always {
            archiveArtifacts artifacts: 'golang-sample-macos-arm64', fingerprint: true
        }
    }
}
