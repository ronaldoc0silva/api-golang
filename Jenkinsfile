pipeline {
    agent any

    stages {
        stage ('Build Image') {
            steps {
                script {
                    dockerapp = docker.build("ronaldosilva00/api-golang:${env.BUILD_ID}", '-f ./Dockerfile .')
                }
            }
        }
    }
}