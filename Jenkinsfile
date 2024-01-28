pipeline {
    agent any

    stages {
        stage ('Inicial') {
            steps {
                script {
                    dockerapp = docker.build("ronaldosilva00/api-golang:v1.2", '-f ./Dockerfile .')
                }
            }
        }
    }
}