pipeline {
    agent any

    environment {
        IMAGE = "alidevops8/hello-devops:latest"
    }

    stages {
        
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                sh """
                docker build -t $IMAGE ./app
                """
            }
        }
        
        stage('Push Docker Image') {
            steps {
                withDockerRegistry(credentialsId: 'dockerhub', url: "") {
                    sh """
                    docker push $IMAGE
                    """
                }
            }
        }

        stage('Notify ArgoCD') {
            steps {
                echo "Trigger ArgoCD will be added later"
            }
        }
    }

    post {
        success {
            echo "🚀 Build complete and pushed to Docker Hub!"
        }
        failure {
            echo "❌ Pipeline failed"
        }
    }
}
