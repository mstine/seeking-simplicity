---
date: '2007-06-27 12:54:00'
layout: post
slug: how-to-implement-form-level-validation-in-jsf
status: publish
title: How to implement form-level validation in JSF
wordpress_id: '36'
categories:
- java
- jsf
- validation
---

Recently I was faced with the challenge of implementing form-level (or page-level) validation in a JSF-based application. What I mean by form-level validation is the need to evaluate a subset of a form's fields as a unit, rather than simply validating each field in isolation. An example of this type of validation can be found on a user registration form where one has to select a password in one text field, and then retype the same password in another text field for confirmation. Validating that these two text fields contain the same password is an example of form level validation.  
  
In my case, I had two date selector components on my form, one for a start date/time and one for an end date/time for an event that was being scheduled. The rule I needed to validate was that the end date/time was later than the start date/time.  
  
There are a few ways to implement validation like this, including but not limited to:  
  


  

  1. Build a custom component that renders selectors for both the start and end date/time, then validate as a unit. This actually is field-level validation and doesn't truly address the form-level problem.
  

  2. Implement a validator method on a managed bean that will evaluate the data submitted for multiple components.
  
  
  
I'll address the second method in this HOWTO.  
  
First, you'll need to bind at least _n-1_ of the components that you want to validate to properties on your managed bean. The simplest way is to declare properties of type **UIInput**:  


>   

>     
>     <br></br>private UIInput startDateComponent;<br></br><br></br>public UIInput getStartDateComponent() {<br></br>    return startDateComponent;<br></br>}<br></br><br></br>public void setStartDateComponent(UIInput startDateComponent) {  <br></br>    this.startDateComponent = startDateComponent;<br></br>}<br></br>
> 
>   


  
and do the actual binding in the JSP:  
  


> <br></br><t:inputDate id="eventStart" value="#{orderForm.sampleInfo.requestedStartTime}"<br></br>    type="both"<br></br>    popupCalendar="true"<br></br>    ampm="true" binding="#{dateValidationForm.startDateComponent}"/><br></br>
> 
>   


  
Next, you'll implement the validation method, which can have any name you like, but must share the same signature as this example:  


> <br></br>public void validateEndDate(FacesContext context, UIComponent toValidate, Object value) {<br></br>    Date endDate = (Date) value;<br></br>    Date startDate = (Date) getStartDateComponent().getLocalValue();<br></br><br></br>    if (startDate == null) {<br></br>        context.addMessage(getStartDateComponent().getClientId(context),new FacesMessage("Please specify a valid date and time."));<br></br>        throw new ValidatorException(new FacesMessage());<br></br>    }<br></br><br></br>    long endTime = endDate.getTime();<br></br>    long startTime = startDate.getTime();<br></br><br></br>    if (startTime >= endTime) {<br></br>        addError("errors.batchOrder.invalidEndDate");<br></br>        throw new ValidatorException(new FacesMessage("Event end must be later than event start."));<br></br>    }<br></br>}<br></br>

  
And finally, you'll bind the validation method to the last component in your subset of components that need to be validated together:  


> <br></br><t:inputDate id="eventEnd" value="#{orderForm.sampleInfo.requestedEndTime}" type="both" popupCalendar="true" ampm="true" validator="#{dateValidationForm.validateEndDate}"/><br></br></blockquote>
> 
>   
To understand why I say _n-1_ components, think of the way the validation phase occurs in JSF. Data is bound to the components in the order that they occur in the JSF component tree, which just so happens to be the order in which they appear in the JSP source. Looking at the **validateEndDate** method, you'll see that I only reference the **startDateComponent** from the binding, but I reference the **endDate** as the **Object value** reference that was passed into the method. This is why you only need to bind _n-1_ components, because you get the _nth_ component from the method signature.  
  
If you want to be more uniform and bind all of the components, you could create an extra dummy hidden value component and bind the validator method to it. You could then bind all of the components to your managed bean and access them all from the bindings rather than accessing one from the method signature.  
  
The **validateEndDate** method itself is rather simple. First you access the data by getting the local value of each component (again, the **endDate** value is not accessed in this way - in fact, it hasn't been bound yet because it must be validated first, and that's what's happening in this method!). You then apply the business rule. You'll see that first I look to see if the **startDate** is null. I'm not sure why this is possible, but if the **startDate** was not submitting a good value on the FIRST submit, the local value was null. So, I catch that here. I add an error message to the **startDateComponent** and throw a **ValidatorException**. If the business rule is violated, throw a **ValidatorException**. (I'm also using the **addError** method provided by AppFuse to work w/ its message framework as well, but that is not necessary w/ all JSF apps).  
  
Now, for the final problem I encountered. In Weblogic server, which we're still using for the time being, if your session cannot be serialized then it deletes your entire session. Obviously this can cause major problems in any web app. To deal with this, ANY SESSION SCOPED MANAGED BEAN must be fully serializable, meaning it and any objects referenced in its state. Herein lies the problem for JSF. Instances of **UIComponent** (an ancestor of **UIInput**) are not serializable, so if we bind our components to **UIInput** fields on a session-scoped managed bean (the bean backing this form is an Order Form/Shopping Cart style bean), it will not be serializable and Weblogic will kick out your session.  
  
To deal with this problem, realize that there is no reason that you can only have one managed bean backing a form. In fact, you can reference as many managed beans as you need. Since validation is done for each request, there is no need to manage any state there across multiple requests like we need to do with a shopping cart. So, why not declare an additional managed bean that is REQUEST SCOPED, and then put the bindings and validation method there. That is exactly what I did. Here is the entire bean:  

>
>> <br></br>public class DateValidationForm extends BasePage {<br></br><br></br>    private UIInput startDateComponent;<br></br><br></br>    public UIInput getStartDateComponent() {<br></br>        return startDateComponent;<br></br>    }<br></br> <br></br>    public void setStartDateComponent(UIInput startDateComponent) {<br></br>        this.startDateComponent = startDateComponent;<br></br>    }<br></br><br></br>    public void validateEndDate(FacesContext context, UIComponent toValidate, Object value) {<br></br>        Date endDate = (Date) value;<br></br>        Date startDate = (Date) getStartDateComponent().getLocalValue();<br></br><br></br>        if (startDate == null) {<br></br>            context.addMessage(getStartDateComponent().getClientId(context),new FacesMessage("Please specify a valid date and time."));<br></br>            throw new ValidatorException(new FacesMessage());<br></br>        }<br></br><br></br>        long endTime = endDate.getTime();<br></br>        long startTime = startDate.getTime();<br></br><br></br>        if (startTime >= endTime) {<br></br>            addError("errors.batchOrder.invalidEndDate");<br></br>            throw new ValidatorException(new FacesMessage("Event end must be later than event start."));<br></br>        }<br></br>    }<br></br>}<br></br>
> 
>   
and the declaration in faces-config.xml:  

>
>> <br></br><managed-bean><br></br>    <managed-bean-name>dateValidationForm</managed-bean-name><br></br>    <managed-bean-class>org.stjude.hc.srmcti.webapp.action.ordering.
>>     DateValidationForm</managed-bean-class><br></br>    <managed-bean-scope>request</managed-bean-scope><br></br></managed-bean><br></br>
> 
>   
The added bonus is that you can reuse this bean across all forms where you need this behavior. My application happens to have 2 additional forms where I would have repeated this logic, so I just reference this bean there.  
  
Enjoy!
