pipeline {
    agent {label 'jenkins_slave'}

    environment {

          DOCKER_HUB_ACCOUNT="shajalahamedcse"
          DOCKER_IMAGE_NAME="webpage"
          K8S_DEPLOYMENT_NAME="webapp"
          PASS=""
       }

    stages {
        stage('Docker Build and push') {
            steps {
                echo 'Building..'
                sh ("sudo docker login -u=${DOCKER_HUB_ACCOUNT} -p=${PASS}")
                sh ("sudo docker build -t ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:${env.BUILD_NUMBER} .")
                sh ("sudo docker push ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:${env.BUILD_NUMBER}")
                sh ("sudo docker tag ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:${env.BUILD_NUMBER} ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:latest")
                sh ("sudo docker push ${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:latest")

            }
        }

        stage("Get Kuber Node"){
            steps {
                sh("pwd")
                sh("kubectl get nodes")
                sh("kubectl set image deployment/${K8S_DEPLOYMENT_NAME} ${K8S_DEPLOYMENT_NAME}=${DOCKER_HUB_ACCOUNT}/${DOCKER_IMAGE_NAME}:${env.BUILD_NUMBER}")
            }
        }
    }

}
