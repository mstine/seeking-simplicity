---
date: '2009-03-13 22:08:00'
layout: post
slug: grailsdojo-ajax-file-upload
status: publish
title: Grails+Dojo Ajax File Upload
wordpress_id: '68'
categories:
- CodeProject
---

I'm currently working on a [Grails](http://grails.org/) application with a requirement for uploading product images, which are subsequently resized into thumbnails of various sizes for different shopping screens. Since it took a bit of doing to get this done, I thought I'd post my solution here in case anyone could benefit from it.

First, to set the stage, I've upgraded the application all the way to the newly released [Grails 1.1](http://www.springsource.org/node/1107).

To start off, I naively attempted to use the built-in Prototype library to do the upload with a simple `<g:submitRemote/>` tag. You might have guessed that this didn't work at all. Continuing to learn the hard way (without reading the manual I might add), I submitted this as a bug to the Grails project. Graeme ever so politely informed me that this was known and expected behavior, as Prototype doesn't support such a thing. However, it was Graeme that also tipped me off to try Dojo.

So, continuing down this path, I proceeded to install the Grails Dojo plugin. Once this is done, a `<g:javascript library="dojo"/&gt` is supposedly all that is required to convert your Ajax calls from Prototype to Dojo. This turned out to not be the case for me, with Javascript errors popping up all over the place, not the least of which was that dojo.js seemed to be installed in an unexpected location via the plugin. The hacker in me simply copied this to the expected location and moved along. However, as I attempted to work with Dojo's file upload support, I discovered that the version of Dojo shipped with the plugin seemed at first glance to be way behind. Frustrated by this, I went ahead and stripped the Dojo plugin out and installed [the latest version available (at this writing 1.2.3)](http://download.dojotoolkit.org/release-1.2.3/dojo.js), and set about to develop my solution.

Here goes:

**1. Add the necessary Dojo dependencies to your GSP:**


    
    
    <script type="text/javascript" src="${createLinkTo(dir: 'js/dojo', file: 'dojo.js')}"
       djConfig="parseOnLoad:true, isDebug:true"></script>
    <g:javascript>
       dojo.require("dojo.io.iframe");
    </g:javascript>
    



**2. Write a function using dojo.io.iframe to send the form:**


    
    
    function submitForm() {
       dojo.io.iframe.send({
          form: 'uploadProductImageForm',
          load: function (data) {
             dojo.byId('productImage').innerHTML = data;
          }
       });
    }
    



**3. Create the file upload form:**


    
    
    <g:form name="uploadProductImageForm" method="post" action="uploadProductImage" enctype="multipart/form-data">
       <input type="hidden" name="id" value="${productInstance?.id}"/>
       <input type="file" name="newProductImage"/>
       <span class="button"><input class="add" type="button" name="uploadImage" value="Upload Image" onclick="submitForm()"/></span>
    </g:form>
    



**4. Create the controller method:**


    
    
    def uploadProductImage = {
       def f = request.getFile('newProductImage')
       if (!f.empty) {
    
       def imagePath = grailsApplication.config.store.productImages.location
    
       //Create unique name for this image set based on current timestamp
       def name = "image" + new Date().getTime()
    
       //Store the file
       def file = new File("${imagePath}/${originalFilename}")
       f.transferTo(file)
    
       //Do some image processing (resizing, etc.)
       ...
    
       //Dojo requires returning the result nested in an HTML document containing a body and textarea tag. Do this with
       //Groovy's built-in MarkupBuilder
    
       def writer = new StringWriter()
       def xml = new MarkupBuilder(writer)
    
       xml.html {
         body {
           textarea {
             img(src: resource(dir: grailsApplication.config.store.productImages.webPath, file: product.mediumImage.name), width: '250')
           }
         }
       }
    
       render writer.toString()
       }
       else {
          flash.message = 'file cannot be empty'
          redirect(action: show)
       }
    }
    



And there you have it. Let me know what you think of this solution. It definitely works for me. You will notice that I didn't include an upload progress bar - I'll be doing this in a future iteration of the project. Cheers!
