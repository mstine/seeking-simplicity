---
date: '2009-06-01 23:56:06'
layout: post
slug: deploying-grails-to-morph-appspace-communityone-2009-lightning-talk
status: publish
title: 'Deploying Grails to Morph AppSpace: #CommunityOne 2009 Lightning Talk'
wordpress_id: '174'
categories:
- cloud
- grails
- javaone2009
---

I gave two lightning talks at [CommunityOne](http://developers.sun.com/events/communityone/2009/west/index.jsp) today, the first of which described deploying [Grails](http://grails.org) applications to [Morph AppSpace](http://mor.ph/products_appspace).

For the uninitiated, Grails is a Ruby on Rails inspired full stack web development framework which brings "convention over configuration" and "DRY" into the Java web development arena. Unlike Rails, it is not an effort from scratch, but rather stands on the shoulders of proven giants in the Java world like the Spring framework and Hibernate. It does this using Groovy, the popular dynamic scripting language for the JVM, as a sort of "DSL for web development." Find it at [http://grails.org](http://grails.org).

Morph AppSpace on the other hand is a fully-configured and managed environment for hosting web applications, and currently supports Java, Grails, Rails, and PHP applications. It is a "platform as a service" (PaaS) provider that abstracts away the details of Amazon EC2 and S3 technologies. Systems architecture, backups, monitoring, failover, scalability - all of this is handled by Morph. You simply develop and deploy your application - Morph does the rest. Find it at [http://mor.ph/products_appspace](http://mor.ph/products_appspace).

So to get going, visit [http://mor.ph](http://mor.ph) and sign up for a free developer account. Create yourself a Java application subscription, and pick your choice of database (MySQL or PostgreSQL). Create the database, and then download two very important files into the root directory of your Grails project: deployment.properties, which contains the metadata describing your application to the Morph AppSpace platform, and morph_deployer.jar, which contains the client API to the platform.

Next you'll need to install the [Grails morph-deploy plugin](http://grails.org/plugin/morph-deploy). If you're using Grails 1.1, you'll need to checkout [the trunk version from SVN](https://svn.codehaus.org/grails-plugins/grails-morph-deploy/trunk/), as the version in the plugin repository is not 1.1 ready. Install this plugin locally by running "grails install-plugin $PATH_TO_PLUGIN." Next, you'll need to edit DataSource.groovy to contain the following:


    
    
    production {
            dataSource {
                driverClassName = 'com.mysql.jdbc.Driver'
                dbCreate = "update"
                jndiName = "java:comp/env/jdbc/morph-ds"
                dialect = 'org.hibernate.dialect.MySQLDialect'
            }
    }
    



Finally, run "grails war" to build the war file, and "grails deploy" to upload your application to the platform. Once the upload is complete, visit the management interface and check the logs to see that you've successfully deployed. Once it's finished, click on the link to your application. Happy Grails on the cloud!

Here's the screencast from my talk. Enjoy!

[youtube=http://www.youtube.com/watch?v=JYPJ26-1YTM&hl;=en&fs;=1&border;=1]
