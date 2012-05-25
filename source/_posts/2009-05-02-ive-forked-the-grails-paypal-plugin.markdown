---
date: '2009-05-02 00:59:46'
layout: post
slug: ive-forked-the-grails-paypal-plugin
status: publish
title: I've "forked" the Grails PayPal Plugin...
wordpress_id: '158'
categories:
- CodeProject
- grails
- groovy
---

I'm currently completing the finishing touches on a new e-commerce site for my wife's stationary business. We decided a long time ago to use [PayPal](http://www.paypal.com) for all of the payment processing since we've had a great experience using it for our eBay selling. About a year ago Graeme Rocher polished off the last release of a [PayPal plugin](http://grails.org/plugin/paypal) that is available in [the Grails Plugin repository](http://grails.org/plugin/home). It really is a very nice plugin, yet I had a couple of problems with it:




	
  * It is only capable of handling payments for one item transactions via "Buy Now" buttons. I want to upload an entire shopping cart full of multiple items.

	
  * It currently won't handle shipping addresses. The site I'm building allows the user to maintain a list of shipping addresses, and I'd want to send the address information they choose along to PayPal.

	
  * Minor issue: needed to upgrade the plugin to Grails 1.1.



So, I decided this evening to fork the plugin. I want to leverage all of the great work that has been done thus far (especially with the IPN processing part - superb stuff), but I have to add in these two functions and do the version upgrade. Interestingly enough I couldn't get the tests to run out of the box after the upgrade. No good developer likes to modify code without a stable running test suite, right? So what I ended up doing was creating a new Grails 1.1 plugin project and copying the original artifacts over. Once this was done all of the tests ran perfectly.

So, at this point I've added function #1. You can now redirect to the "uploadCart" action. It assumes that you have already constructed a Payment object (now containing PaymentItems) and saved it, and then passed the transactionId along. I did this so that folks with multiple different ways of handling shopping carts could have some degree of flexibility - the onus is on you to map your cart to your Payment object correctly. The original functions implemented by Graeme are backward compatible, assuming only one PaymentItem in the Payment. I've run all of the original tests and also did some manual functional testing by running the plugin app against my own PayPal Sandbox account. So far so good. Look for more updates as this evolves. Once I get something I'm totally happy with I'll see about getting it pushed back into the main plugin repo.

Want to take a closer look? Visit [http://github.com/mstine/grails-paypal-plugin/tree/master](http://github.com/mstine/grails-paypal-plugin/tree/master)
