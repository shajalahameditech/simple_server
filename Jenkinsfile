pipeline {
    agent {label 'jenkins_slave'}

    stages {
        stage('Build') {
            steps {
                echo 'Building..'

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