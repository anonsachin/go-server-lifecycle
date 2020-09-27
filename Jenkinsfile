pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'docker build -t server . `
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'

                sh 'docker run --name server --network host server'
            }
        }
    }
}