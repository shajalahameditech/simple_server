node {
    def app

    stage('Clone repository') {
        /* Let's make sure we have the repository cloned to our workspace */

        checkout scm
    }

    stage('Build image') {
        /* This builds the actual image; synonymous to
         * docker build on the command line */

        app = docker.build("shajalahamedcse/webpage")
    }

    stage('Test image') {
        /* Ideally, we would run a test framework against our image.
         * Just an example */

        app.inside {
            sh 'echo "Tests passed"'
        }
    }

    stage('Push image') {
        /* Finally, we'll push the image with two tags:
         * First, the incremental build number from Jenkins
         * Second, the 'latest' tag.
         * Pushing multiple tags is cheap, as all the layers are reused. */
        docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
            app.push("${env.BUILD_NUMBER}")
            app.push("latest")
        }
    }

        stage('Send Rocket Chat Message') {
            sh "curl -i -H \"X-Auth-Token: b-EYSzcBzvNjuJe02wOSp-n9fNE2ecYUioVro0Weqy6\" -H \"X-User-Id: DdkKdpdc5siNmJap5\" -H \"Content-type:application/json\" https://chat.internal.ypto.space/api/v1/chat.postMessage  -d '{ \"channel\": \"#general\", \"text\": \"Test bot for infra\" }'"
        }
}
