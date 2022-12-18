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
    }
}
