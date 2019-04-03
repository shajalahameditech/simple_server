pipeline {
    agent {label 'jenkins_slave'}

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh "kubectl get nodes"
                echo 'This will run only if successful'

            }
        }
    }

    post {
        success {
            sh "kubectl get nodes"
            echo 'This will run only if successful'
        }
    }
}