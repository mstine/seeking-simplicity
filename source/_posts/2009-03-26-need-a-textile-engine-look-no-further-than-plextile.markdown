---
date: '2009-03-26 16:23:46'
layout: post
slug: need-a-textile-engine-look-no-further-than-plextile
status: publish
title: Need a Textile engine? Look no further than Plextile!
wordpress_id: '81'
categories:
- CodeProject
- grails
- groovy
- textile
---

In finishing up the first release of JUG Nexus, the open source engine ([http://github.com/mstine/jug-nexus/tree/master](http://github.com/mstine/jug-nexus/tree/master)) behind the new Memphis JUG website, I needed to put a good [Textile](http://www.textism.com/tools/textile/) engine in place. I really don't care for writing actual HTML in a content management system, and JUG Nexus being a lightweight CMS, I wanted a lightweight markup syntax for entering the details of upcoming JUG meetings. Textile is exactly that, and is also used for entering content into the very useful [JUGEvents](http://jugevents.org) system produced by [JUG Padova](http://www.jugpadova.it/) for the Java User Group community at large. I tried out several different engines, and none of them seemed to be very robust when it came to edge cases in the markup. For example, if I want to produce a hyperlink to a website, the textile format is the following:

`
"Linked Text":http://www.site.com
`

This syntax will result in the following HTML:


    
    
    <a href="http://www.site.com">Linked Text</a>
    



Unfortunately, if you have a bang (!) in your "Linked Text," most of the Textile engines that I used would not recognize it as an HTML link. Plextile does!

The only drawback, however minimal, to using Plextile is that it does not come with a pre-packaged JAR file. You have to take the compiled code and JAR it yourself. Now, for free, I'll include how I integrated Plextile with Grails. Grails has a very nice codec feature that provides a facility to register encoders and decoders of textual data as methods on any object. Grails searches for classes following the convention `XCodec` and dynamically registers `encodeAsX` and `decodeX` methods on `java.lang.Object` so that any data can be encoded and decoded. Enter the `TextileCodec`:


    
    
    import com.plink.plextile.TextParser
    
    class TextileCodec {
    
       static encode = {str ->
          new TextParser().parseTextile(str, true)
       }
    
    }
    



Believe it or not, that's it! Here's the GSP template for rendering a JUG meeting:


    
    
    <div class="post">
      <h2>${event.title}<br/><br/><g:formatDate format="EEEE, MMMM dd, yyyy" date="${event.startTime}"/><br/><g:formatDate format="h:mm" date="${event.startTime}"/>-<g:formatDate format="h:mm a" date="${event.endTime}"/><br/><g:if test="${!event.archived}"><a href="http://jugevents.org/jugevents/event/${event.jugEventsId}">Click HERE to register!</a></g:if></h2>
      <div class="entry">
        <h2 class="title">Speaker/Topic:</h2>
        <p>${event.description.encodeAsTextile()}</p>
        <h2 class="title">Location/Directions:</h2>
        <p>${event.location}</p>
        <p>${event.directions.encodeAsTextile()}</p>
      </div>
    </div>
    



Enjoy!
