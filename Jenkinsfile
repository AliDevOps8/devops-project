pipeline {
    agent any

    environment {
        IMAGE_NAME = "alidevops8/hello-devops"
        IMAGE_TAG = "${env.GIT_COMMIT.take(7)}"
        IMAGE = "${IMAGE_NAME}:${IMAGE_TAG}"
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

        stage('Update GitOps Repo') {
            steps {
                sshagent(credentials: ['github-token']) {
                    sh """
                    git clone git@github.com:AliDevOps8/devops-project.git
                    sed -i '' "s%${IMAGE_NAME}:.*%${IMAGE_NAME}:${IMAGE_TAG}%g" ./k8s/deployment.yaml
                    git config user.email "ali.devops8@gmail.com"
                    git config user.name "alidevops8"
                    git commit -am "Update image version to ${IMAGE_TAG}"
                    git push -u origin main
                    """
                }
            }
        }

        stage('Notify ArgoCD (optional)') {
            steps {
                echo "ArgoCD will detect Git changes automatically!"
            }
        }
    }

    post {
        success {
            echo "🚀 CI/CD + GitOps pipeline completed. ArgoCD will sync soon."
        }
        failure {
            echo "❌ Pipeline failed!"
        }
    }
}
