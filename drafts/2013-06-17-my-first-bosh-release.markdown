---
layout: post
title: "My First BOSH Release"
date: 2013-06-17 23:10
comments: true
categories: 
---

This is going to start as a working draft, as I work my way through the steps. I'll clean it up later.

Part I: Develop the release, test w/ Nise BOSH
Part II: Develop the deployment manifest, deploy w/ Micro BOSH

Let's use Nise BOSH to help develop the release in concert with bosh\_cli. TODO: Intro these...

First things first:

``` bash
$ bosh init release tomcat
Release directory initialized
```

This creates us a nice little scaffolded directory structure:

``` bash
$ tree tomcat
tomcat
├── blobs
├── config
│   └── blobs.yml
├── jobs
├── packages
└── src

5 directories, 1 file
```

Let's cover what each of these represents: (TODO: write up the descriptions)

* blobs
* config/blobs.yml
* jobs
* packages
* src

We'll need two tarballs to get a Tomcat instance up and running: a Java(tm) Runtime Environment and Apache Tomcat. Let's retrieve those.

We'll place each of them in a folder beneath blobs.

``` bash
$ cd tomcat
$ mkdir blobs/tomcat
$ curl -O http://mirror.cogentco.com/pub/apache/tomcat/tomcat-7/v7.0.41/bin/apache-tomcat-7.0.41.tar.gz

...curl output omitted

$ cd ..
$ mkdir blobs/java
$ curl -O http://sdlc-esd.sun.com/ESD6/JSCDL/jdk/7u21-b11/jre-7u21-linux-x64.tar.gz?AuthParam=1371536753_2dfcf7cabf130c87f1011d35357e6382&GroupName=JSC&FilePath=/ESD6/JSCDL/jdk/7u21-b11/jre-7u21-linux-x64.tar.gz&File=jre-7u21-linux-x64.tar.gz&BHost=javadl.sun.com

...curl output omitted

mv jre-7u21-linux-x64.tar.gz\&File=jre-7u21-linux-x64.tar.gz\&BHost=javadl.sun.com jre-7u21-linux-x64.tar.gz
```

Yeah...that JRE link is ugly. Gotta love Oracle. Also, TODO: how to get curl to download w/ the right file name?

So, why are we placing these in blobs rather than source? Well, because both of these represent binary releases of the software. I don't need to compile them. (TODO: Does BOSH *really* care about the distinction?)

## Packages

So, the next step is we want to generate packages. Tell more about packages.

``` bash
$ cd ~/Projects/bosh/release_dev/tomcat
$ bosh generate package java
create  packages/java
create  packages/java/packaging
create  packages/java/pre_packaging
create  packages/java/spec

Generated skeleton for `java' package in `packages/java'
```

Neat! More scaffolding. Let's touch on what each of these do:

* packaging
* pre\_packaging
* spec

We'll need to do the same for tomcat:

``` bash
$ bosh generate package tomcat

...output omitted
```

We end up with the same structure now for Tomcat. Time to crack open these files and see what's necessary.

We're actually going to remove the pre\_packaging file from both of our packages, as we won't need its contribution. Let's take a look at the packaging script for the JRE:

``` bash
# abort script on any command that exit with a non zero value
set -e

tar xzf java/jre-7u21-linux-x64.tar.gz
cp -a * ${BOSH_INSTALL_TARGET}
``` 

Note that we're working from an MVP perspective...this will get Java installed, but we may want to do more.

Now we need to look at the spec:

``` yaml
---
name: java
dependencies: []
files:
- java/jre-7u21-linux-x64.tar.gz
```

Again, TODO an overview of what this stuff actually means.

We'll create similar files in the Tomcat package that we created:

``` bash
# abort script on any command that exit with a non zero value
set -e

tar xzf tomcat/apache-tomcat-7.0.41.tar.gz
cp -a * ${BOSH_INSTALL_TARGET}
```

``` yaml
---
name: tomcat
dependencies: []
files:
- tomcat/apache-tomcat-7.0.41.tar.gz
```

## Jobs

OK, now on to jobs. What are jobs?

We only need a job for Tomcat, as that is the "running entity" that we are trying to release. Let's generate the scaffolding:

``` bash
$ bosh generate job tomcat
create  jobs/tomcat
create  jobs/tomcat/templates
create  jobs/tomcat/spec
create  jobs/tomcat/monit

Generated skeleton for `tomcat' job in `jobs/tomcat'
```

Once again, what is this stuff?

* templates
* spec
* monit


