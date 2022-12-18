pipeline {
    agent {
        docker { image 'golang:alpine3.16' }
    }
    stages {
        stage('Test') {
            steps {
                sh 'go version'
            }
        }
        stage('Build') {
            steps {
                sh 'GOOS=darwin GOARCH=arm64 go build -o golang-sample-macos-arm64 main.go'   
            }
        }
    }
    post {
        always {
            archiveArtifacts artifacts: 'golang-sample-macos-arm64', fingerprint: true
        }
    }
}
