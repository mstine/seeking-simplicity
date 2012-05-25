---
date: '2008-02-04 17:43:00'
layout: post
slug: a-javascript-code-prettifier
status: publish
title: A Javascript Code "Prettifier"
wordpress_id: '58'
---

I came across a [link](http://code.google.com/p/google-code-prettify/) today that would allow you to post "pretty" code in any HTML page. I've had plenty of problems (as have others) posting well-formatted source code to my Blogger-based blog. I've wanted to correct this problem for some time, and this project appears to be a decent solution. Here's an example:  
  

    
    function chooseWinner() {<br></br>   $('p3').removeClassName('winner');<br></br>   participantBo.chooseWinnerForEvent(eventId, function(data) {<br></br>      people = $A(data);<br></br>      updateDivs();<br></br>      intervalId = setInterval(doShift, 500);<br></br>   });<br></br>}

  
  
Now that isn't too bad is it? There seems to be a bit of a delay as the Javascript is loaded, but I think that's more a symptom of my connection than it is a problem with the script.  
  
Here's how you can take advantage of this solution (quoted from [http://google-code-prettify.googlecode.com/svn/trunk/README.html](http://google-code-prettify.googlecode.com/svn/trunk/README.html)):  
  


  
     
  1. [Download](http://google-code-prettify.googlecode.com/files/prettify-small.zip) a distribution  
     
  2. Include the script and stylesheets in your document  
       (you will need to make sure the css and js file are on your server, and  
        adjust the paths in the script and link tag)  
       
    
    <link href="prettify.css" type="text/css" rel="stylesheet" /><br></br><script type="text/javascript" src="prettify.js"></script>

  
     
  3. Add `onload="prettyPrint()"` to your  
     document's body tag.  
     
  4. Modify the stylesheet to get the coloring you prefer
  
     
  5. Put code snippets in  
   <pre class="prettyprint">...</pre>  
   or <code class="prettyprint">...</code>  
   and it will automatically be pretty printed.
  
   
