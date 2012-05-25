---
date: '2009-11-10 13:50:11'
layout: post
slug: securing-grails-plugin-artifacts-with-filters
status: publish
title: Securing Grails Plugin Artifacts with Filters
wordpress_id: '278'
categories:
- grails
- groovy
---

So you've just installed the handy dandy Spring Security plugin (http://grails.org/plugin/acegi), which makes it incredibly easy to secure entire Grails controllers and/or controller actions with annotations, such as the following:



This is enabled by turning on controller annotations in your SecurityConfig.groovy file:



So all is now good in our project. We can secure either controllers or actions with annotations, enabling us to declaratively setup security side-by-side with the code that we're securing in a very straightforward manner. You can continue developing your Grails applications with glee, fully assured that security is no longer an issue. But wait, one day you decide to install one of the many useful Grails plugins that add controller artifacts to your application. Lo and behold, you have no way to secure those controllers! Of course, you could descend into $USER_HOME/.grails/$GRAILS_VERSION/projects/projectName/plugins/pluginX and hack the source code for your individual instance of the plugin. This ought to work, but you're now rather constrained in that every time you update the plugin you'll need to remember to go make this manual change. That doesn't sound very agile at all, does it? OK, so how about forking the plugin? This is a little bit better, but now you have the burden of merging changes from the global plugin repository to yours every time a new release happens. This is better, but still a bit cumbersome. How about becoming a committer and adding it to the global source? Of course not. Not everyone will want to secure their plugins the same way you do, and you've just introduced a rather unnecessary dependency on the Spring Security plugin. I say all this in an attempt to paint a grim picture. In reality, we're actually in very good shape. Grails Filters to the rescue!

All that you need to do is create a Grails filter that will match requests to the plugin artifact in question and then delegate to Spring Security for authorization. If they are authorized, you simply return true. If not, you can direct them to your login screen. It's this simple:



As you can see here, I've secured both the Blurb plugin and the Settings plugin in this manner by requiring that the logged in user be in the ROLE_ADMIN role. Now as Glen Smith would say, that's a snack!

_**Update:** [Burt Beckwith](http://burtbeckwith.com/blog/) enlightened me to an approach that will get this done without the use of filters that will also direct you to the requested URL after login rather than the main page. Unfortunately I've never been able to track this down before. Just add the following to SecurityConfig.groovy:_


