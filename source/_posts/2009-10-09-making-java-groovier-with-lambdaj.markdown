---
date: '2009-10-09 09:37:06'
layout: post
slug: making-java-groovier-with-lambdaj
status: publish
title: Making Java "Groovier" with LambdaJ
wordpress_id: '257'
categories:
- functional-programming
- groovy
- java
- loty
---

I spent the better part of yesterday tracing my way through the codebase for a large-scale enterprise application that my team is building right now, and I happened upon the following piece of code:

``` java
//...imports excluded for clarity
public class BusinessActivityBinMetaClassHelper {
//...
   public static List<Long> getSrmMetaClassIdListJava(List<BusinessActivityBinMetaClass> businessActivityBinMetaClassList) {
      List<Long> srmMetaClassIdList = new ArrayList<Long?();

      if(businessActivityBinMetaClassList != null) {
         for(BusinessActivityBinMetaClass businessActivityBinMetaClass : businessActivityBinMetaClassList) {
             if(businessActivityBinMetaClass.getSrmMetaClass() != null && businessActivityBinMetaClass.getSrmMetaClass().getSrmMetaClassId() != null) {
               srmMetaClassIdList.add
                  (businessActivityBinMetaClass.getSrmMetaClass().getSrmMetaClassId());
            }
         }
      }

      return srmMetaClassIdList;
   }
//...
}
```

As I read this code, I thought "This just SCREAMS a need for Groovy's closure iterators." So, last night I quickly hacked out fully-equivalent Groovy version of the code:

{% codeblock lang:java %}
class GroovyExample {

   static def getSrmMetaClassIdListGroovy(def businessActivityBinMetaClassList) {
      businessActivityBinMetaClassList?.collect { it?.srmMetaClass?.srmMetaClassId }.findAll { it != null } ?: new ArrayList<Long>();
   }
}
{% endcodeblock %}

Whew! Much nicer. What did we get out of this? Well...




   
  * Groovy's dynamic typing cleaned up all of the unnecessary type declarations - the only static typing is where we return an empty ArrayList of Longs if the argument to our method is null (a bit of odd behavior, but required to make this code equivalent to the parent Java code.

   
  * We were saved 4 painful null checks by the use of Groovy's null safe dereferencing operator (?.) and the simplified ternary operator otherwise affectionately known as the "Elvis Operator" (?:).

   
  * Using Groovy's collect method, we were able to transform the original list into a list containing Long values by passing in a closure that pulls out the property value of interest.

   
  * Because we're using null safe dereferencing, we are actually inserting nulls into our list if any of the dereferencing fails. Therefore, Groovy's findAll Collection filtering method comes to the rescue. We simply provide it with a closure that returns true for all of the values we want to keep - in this case, "it != null."

   
  * Perhaps most importantly, we've shorted our code from 11 lines (including whitespace lines for clarity) to ONE LINE that much more clearly expresses the intent of the code.



Great - why don't we just rewrite the entire application is Groovy? Well, hold on just a minute. At the time we started this application, as much as some of us loved Groovy, we just didn't have enough Groovy mindshare to go there yet. On top of that, we were already experimenting with several new architectural ideas and technologies, and Groovy would have added yet one more risk to the puzzle. I say all this to acknowledge that sometimes you just can't move to another language for one reason or another, regardless of how attractive its features may be.

But let's take a queue from the _Pragmatic Programmer_ and explore the LOTY (Language of the Year) concept one more time. One of the reasons that you're encouraged to learn new languages is to change the way you program in your main language. You may learn Groovy, Scala, Clojure, Ruby, etc., etc., etc. and never use them in your day job - and yet the experience of coding in a new language with new constructs and idioms will necessarily change the way you THINK about programming in every other language.

So, let's think about the possibility of coding something that is much more similar to the Groovy version and yet stick with regular Java code. Fortunately, there are several libraries out there that bring much of the flavor and power of Groovy's closure iterators to Java. I'd like to focus in on one of them, LambdaJ ([http://code.google.com/p/lambdaj/](http://code.google.com/p/lambdaj/)).

LambdaJ provides constructs that allow us to "...manipulate collections in a pseudo-functional and statically typed way." Let's take a look at this example implementing using LambdaJ:

``` java
//...some imports excluded for clarity
import ch.lambdaj.function.convert.Converter
import static ch.lambdaj.Lambda.*
import static org.hamcrest.Matchers.*

public class BusinessActivityBinMetaClassHelper {
//...
   public static List<Long> getSrmMetaClassIdListJava(List<BusinessActivityBinMetaClass> businessActivityBinMetaClassList) {
      return (businessActivityBinMetaClassList != null) ? filter(notNullValue(),convert(businessActivityBinMetaClassList, new IdExtractor())) : new ArrayList<Long>();
   }

   class IdExtractor implements Converter<BusinessActivityBinMetaClass,Long> {
	Long convert(BusinessActivityBinMetaClass from) {
	   if (from.getSrmMetaClass() != null && from.getSrmMetaClass().getSrmMetaClassId() != null) {
	      return from.getSrmMetaClass().getSrmMetaClassId();
	   }
   }
}
```

Not quite as nice as the Groovy code - we got A LOT out of the null-safe dereference and Elvis operators. However, our overall intent is still a bit clearer. Let's analyze:




   
  * First we need to implement one of LambdaJ's Converters. A Converter is nothing more than a generic Interface that defines one method: T convert(F from), where F is the type we're converting from and T is the type we're converting to. In this case, we want to convert an object of type BusinessActivityBinMetaClass to an object of type Long. Our implementation determines how this conversion takes place, in this case by extracting the id property from its child.

   
  * Next, after statically importing the methods of ch.lambdaj.Lambda, we call the convert method, passing it our List and our newly implemented Converter. This gives us the equivalent of Groovy's collect method, with the Converter taking the place of the closure.

   
  * We're still shoving nulls into our List with this implementation, so we further filter our list using LambdaJ's "filter" method, passing it the list returned by "filter," and a [Hamcrest](http://code.google.com/p/hamcrest/) matcher that returns only notNullValue()'s.

   
  * Finally, we use our old friend the ternary operator to return the empty ArrayList of Long values if our method argument is null.



If you don't count the Converter implementation, we've gotten ourselves down to 2 lines of code (1 if you don't mind long lines). In this case I implemented IdExtractor as a named inner class - you could do this with an anonymous inner class and it would look a lot more like a closure, but the added noise of all of the null checking made the undesirable for me. Perhaps if your code has less noise (or guarantees that values aren't null), that would be a better approach.

Another alternative is to make IdExtractor a top-level class that, if general enough, could be reused throughout the codebase. The benefit of this is that you now have a nice code unit rather than a battery of static methods in a utility class, and unit testing becomes much more clean and elegant.

So, we've still made some progress and made our code a bit more Groovy. I encourage you to explore LambdaJ's feature set and see how it might make your code a bit more concise with clearer intent. And just to stir up a little controversy, look what would have happened in Java 7 had the null safe dereference and Elvis operator's made [the Project Coin cut](http://blogs.sun.com/darcy/entry/project_coin_final_five):

``` java
//...some imports excluded for clarity
import ch.lambdaj.function.convert.Converter
import static ch.lambdaj.Lambda.*
import static org.hamcrest.Matchers.*

public class BusinessActivityBinMetaClassHelper {
//...
   public static List<Long> getSrmMetaClassIdListLambdaJ(List<BusinessActivityBinMetaClass> businessActivityBinMetaClassList) {
      return filter(notNullValue(),convert(businessActivityBinMetaClassList,
         new Converter<BusinessActivityBinMetaClass,Long> {
            Long apply(BusinessActivityBinMetaClass from) { return from?.getSrmMetaClass()?.getSrmMetaClassId()}
         })) ?: new ArrayList<Long>();
   }
}
```

Nice, huh? ;-)
