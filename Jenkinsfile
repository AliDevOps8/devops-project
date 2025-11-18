pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh 'docker build -t ghcr.io/alidevops8/hello-devops:latest .'
            }
        }

        stage('Push') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'ghcr', usernameVariable: 'USR', passwordVariable: 'PWD')]) {
                    sh 'echo $PWD | docker login ghcr.io -u $USR --password-stdin'
                }
                sh 'docker push ghcr.io/alidevops8/hello-devops:latest'
            }
        }

        stage('GitOps Sync') {
            steps {
                sh 'echo "Trigger ArgoCD sync (optional webhook)"'
            }
        }
    }
}

