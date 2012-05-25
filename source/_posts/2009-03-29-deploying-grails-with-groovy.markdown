---
date: '2009-03-29 03:38:09'
layout: post
slug: deploying-grails-with-groovy
status: publish
title: Deploying Grails with Groovy
wordpress_id: '89'
categories:
- CodeProject
- deployment
- grails
- groovy
---

Interesting title, eh? Maybe this one will make it through Glen's filter at [GroovyBlogs.org](http://groovyblogs.org).

On to the meat. I've been steadily working on a couple of Grails applications, one being the website for the [Memphis JUG](http://www.memphisjug.org), and another being the e-commerce site for my wife's soon to be launched designer stationery business. Just like your average Grails developer, I've been happily coding away at 127.0.0.1 using the good old development Jetty+MySQL stack. Well, in the last week or so it's come time to actually deploy both of these applications into production. I started out last weekend with the Memphis JUG site. My first approach was to build the WAR file locally and then "scp" it up to the server.

YMMV, but the upload speed on my DSL connection is horrible! After doing this three or four times in one night, waiting 15-20 minutes for the WAR to upload each time (Grails WAR's are rather thick when carrying all of the dependencies along), I decided to myself, "There must be a better way to do this."

Fast forward a week and here I sit working on the first "pre-production" release of my wife's store site. With quite a bit of time on my hands during these "dark and early" hours (a story for a later entry), I decided it was time for the experiment.

Each of these projects is hosted at GitHub, so the process that I mapped out in my mind looked like this:



	
  1. Check out the latest code from GitHub

        
  2. Run "grails war"

        
  3. Stop the Tomcat service (my hosting provider sets up Tomcat to run as a service)

        
  4. Delete the remnants of the previous deployment from Tomcat's deployment directory

        
  5. Copy the new WAR file to Tomcat's deployment directory

        
  6. Start the Tomcat service



By the way, I forgot to mention that before doing all of this I moved the production data source definition from being locally defined to being a JNDI lookup within Tomcat. This posed its own challenge, which I'll blog about a bit later.

Anyway, back to the deployment. I though this would be an excellent opportunity to take [Groovy's Antbuilder](http://groovy.codehaus.org/Using+Ant+from+Groovy) out for a spin. Here's an example of what I put together:


    
    
    #!/usr/bin/env groovy
    
    def ant = new AntBuilder()
    
    //Update the codebase from GitHub
    ant.exec(executable:'git', dir: "${PROJECT_DIR}") {
            arg(value:'pull')
    }
    
    //Generate the WAR file using Ant
    ant.ant(dir:'"${PROJECT_DIR}"', target: 'war')
    
    //Stop Tomcat
    ant.exec(executable:'service') {
            arg(line:'tomcat6 stop')
    }
    
    //Delete the old webapp contents from Tomcat's deploy directory
    ant.delete(includeemptydirs:'true', verbose:'true') {
            fileset(dir:"${CONTEXT_ROOT_DIR}", includes:'**/*')
    }
    
    //Copy the new WAR file to Tomcat's deploy directory
    ant.copy(file:"${WAR_FILE}", tofile:"${CONTEXT_ROOT_DIR}/ROOT.war")
    
    //Start Tomcat
    ant.exec(executable:'service') {
            arg(line:'tomcat6 start')
    }
    



As you can see, I have a few undefined Groovy constants in there. Let's just say I didn't want to expose ALL of the details of my server.

At any rate, it's pretty simple. One word of warning - if you're using Ehcache, make sure to add it to your ivy.xml dependencies, or the Ant build won't bring it in like running "grails war" will. I hope someone finds this simple script useful. Enjoy!
