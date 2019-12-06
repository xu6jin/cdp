pipeline {
  agent any
  stages {
    stage('init') {
      steps {
        tool 'jdk-11'
        tool 'maven-3'
      }
    }

    stage('build') {
      steps {
        sh '''java -version
mvn  -version'''
      }
    }

    stage('archive') {
      steps {
        archiveArtifacts 'user-api-server/target/user-api-server-0.1.jar,mine-api-server/target/mine-api-server-0.1.jar,activity-api-server/target/activity-api-server-0.1.jar,im-api-server/target/im-api-server-0.1.jar,notification-api-server/target/notification-api-server-0.1.jar,blackhole-api-server/target/blackhole-api-server-0.1.jar'
      }
    }

  }
}