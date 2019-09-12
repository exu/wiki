# JAVA programming - [back to index](/)

GO/PHP/JS developer notes on Java.

## Naming converntions

### Interfaces / Class naming

WTF with that all `SomethingImpl` or `ISomething`?

<https://stackoverflow.com/questions/2814805/java-interfaces-implementation-naming-convention>

> If all you can come up with to make your Class name unique is **suffixing it with Impl, then you need to rethink having an Interface at all**. So when you have a situation where you
have an Interface and a single Implementation that is not uniquely specialized from the Interface you probably don't need the Interface.

## Stdlib

### Project Organisation

### Args

### Exceptions

## File handling

## Streaming (Java8 streams)

When our Java project wants to perform the following operations, it’s better to use Java 8

Stream API to get lot of benefits:

- When we want perform Database like Operations. For instance, we want perform groupby operation, orderby operation etc.
- When want to Perform operations Lazily.
- When we want to write Functional Style programming.
- When we want to perform Parallel Operations.
- When want to use Internal Iteration
- When we want to perform Pipelining operations.
- When we want to achieve better performance.

## Spliterator

Spliterator stands for Splitable Iterator. It is newly introduced by Oracle Corporation as
part Java SE 8.  Like Iterator and ListIterator, It is also one of the Iterator interface.

### Differences between Iterator and Spliterator in Java SE 8?

<table border="2" cellspacing="0" cellpadding="6" rules="groups" frame="hsides">

<colgroup>
<col  class="left" />

<col  class="left" />
</colgroup>
<tbody>
<tr>
<td class="left">SPLITERATOR</td>
<td class="left">ITERATOR</td>
</tr>


<tr>
<td class="left">It is introduced in Java SE 8.</td>
<td class="left">It is available since Java 1.2.</td>
</tr>


<tr>
<td class="left">Splitable Iterator</td>
<td class="left">Non-Splitable Iterator</td>
</tr>


<tr>
<td class="left">It is used in Stream API.</td>
<td class="left">It is used for Collection API.</td>
</tr>


<tr>
<td class="left">It uses Internal Iteration concept to iterate Streams.</td>
<td class="left">It uses External Iteration concept to iterate Collections.</td>
</tr>


<tr>
<td class="left">We can use Spliterator to iterate Streams in Parallel and Sequential order.</td>
<td class="left">We can use Iterator to iterate Collections only in Sequential order.</td>
</tr>


<tr>
<td class="left">We can get Spliterator by calling spliterator() method on Stream Object.</td>
<td class="left">We can get Iterator by calling iterator() method on Collection Object.</td>
</tr>


<tr>
<td class="left">Important Method: tryAdvance()</td>
<td class="left">Important Methods: next(), hasNext()</td>
</tr>
</tbody>
</table>

## Optional

What is Optional in Java 8? What is the use of Optional? Advantages of Java 8 Optional?

Optional is a final Class introduced as part of Java SE 8. It is defined in java.util
package.

It is used to represent optional values that is either exist or not exist. It can contain
either one value or zero value. If it contains a value, we can get it. Otherwise, we get
nothing.

It is a bounded collection that is it contains at most one element only. It is an
alternative to “null” value.

Main Advantage of Optional is:

- It is used to avoid null checks.
- It is used to avoid NullPointerException.

## Functional interface

### What is a Functional Interface? What is SAM Interface?

A Functional Interface is an interface, which contains one and only one abstract
method. Functional Interface is also know as SAM Interface because it contains only one
abstract method.  SAM Interface stands for Single Abstract Method Interface. Java SE 8 API
has defined many Functional Interfaces.

### Is is possible to define our own Functional Interface? What is @FunctionalInterface? What are the rules to define a Functional Interface?

Yes, it is possible to define our own Functional Interfaces. We use Java SE 8’s
@FunctionalInterface annotation to mark an interface as Functional Interface.  We need to
follow these rules to define a Functional Interface:

- Define an interface with one and only one abstract method.
- We cannot define more than one abstract method.
- Use @FunctionalInterface annotation in interface definition.
- We can define any number of other methods like Default methods, Static methods.
- If we override java.lang.Object class’s method as an abstract method, which does not
    count as an abstract method.

### Is @FunctionalInterface annotation mandatory to define a Functional Interface? What is the use of @FunctionalInterface annotation? Why do we need Functional Interfaces in Java?

It is not mandatory to define a Functional Interface with @FunctionalInterface
annotation. If we don’t want, We can omit this annotation. However, if we use it in
Functional Interface definition, Java Compiler forces to use one and only one abstract
method inside that interface.  Why do we need Functional Interfaces? The type of a Java SE
8’s Lambda Expression is a Functional Interface. Whereever we use Lambda Expressions that
means we are using Functional Interfaces.

### Runnable example

If we look into some other programming languages such as C++, JavaScript; they are called
functional programming language because we can write functions and use them when
required. Some of these languages support Object Oriented Programming as well as
Functional Programming.

Being object oriented is not bad, but it brings a lot of verbosity to the program. For
example, let’s say we have to create an instance of Runnable. Usually we do it using
anonymous classes like below.

    Runnable r = new Runnable(){
                @Override
                public void run() {
                    System.out.println("My Runnable");
                }};

If you look at the above code, the actual part that is of use is the code inside run()
method. Rest all of the code is because of the way java programs are structured.

Java 8 Functional Interfaces and Lambda Expressions help us in writing smaller and cleaner
code by removing a lot of boiler-plate code.

### Java 8 Functional Interface

An interface with exactly one abstract method is called Functional
Interface. @FunctionalInterface annotation is added so that we can mark an interface as
functional interface.

It is not mandatory to use it, but it’s best practice to use it with functional interfaces
to avoid addition of extra methods accidentally. If the interface is annotated with
@FunctionalInterface annotation and we try to have more than one abstract method, it
throws compiler error.

The major benefit of java 8 functional interfaces is that we can use lambda expressions to
instantiate them and avoid using bulky anonymous class implementation.

Java 8 Collections API has been rewritten and new Stream API is introduced that uses a lot
of functional interfaces. Java 8 has defined a lot of functional interfaces in
java.util.function package. Some of the useful java 8 functional interfaces are `Consumer`,
`Supplier`, `Function` and `Predicate`.

You can find more detail about them in Java 8 Stream Example.

`java.lang.Runnable` is a great example of functional interface with single abstract method
run().

Below code snippet provides some guidance for functional interfaces:

    interface Foo { boolean equals(Object obj); }
    // Not functional because equals is already an implicit member (Object class)

    interface Comparator<T> {
     boolean equals(Object obj);
     int compare(T o1, T o2);
    }
    // Functional because Comparator has only one abstract non-Object method

    interface Foo {
      int m();
      Object clone();
    }
    // Not functional because method Object.clone is not public

    interface X { int m(Iterable<String> arg); }
    interface Y { int m(Iterable<String> arg); }
    interface Z extends X, Y {}
    // Functional: two methods, but they have the same signature

    interface X { Iterable m(Iterable<String> arg); }
    interface Y { Iterable<String> m(Iterable arg); }
    interface Z extends X, Y {}
    // Functional: Y.m is a subsignature & return-type-substitutable

    interface X { int m(Iterable<String> arg); }
    interface Y { int m(Iterable<Integer> arg); }
    interface Z extends X, Y {}
    // Not functional: No method has a subsignature of all abstract methods

    interface X { int m(Iterable<String> arg, Class c); }
    interface Y { int m(Iterable arg, Class<?> c); }
    interface Z extends X, Y {}
    // Not functional: No method has a subsignature of all abstract methods

    interface X { long m(); }
    interface Y { int m(); }
    interface Z extends X, Y {}
    // Compiler error: no method is return type substitutable

    interface Foo<T> { void m(T arg); }
    interface Bar<T> { void m(T arg); }
    interface FooBar<X, Y> extends Foo<X>, Bar<Y> {}
    // Compiler error: different signatures, same erasure

### Lambda Expressions

Lambda Expression are the way through which we can visualize functional programming in the
java object oriented world. Objects are the base of java programming language and we can
never have a function without an Object, that’s why Java language provide support for
using lambda expressions only with functional interfaces.

Since there is only one abstract function in the functional interfaces, there is no
confusion in applying the lambda expression to the method. Lambda Expressions syntax is
(argument) -> (body). Now let’s see how we can write above anonymous Runnable using lambda
expression.

    Runnable r1 = () -> System.out.println("My Runnable");

Let’s try to understand what is happening in the lambda expression above.

Runnable is a functional interface, that’s why we can use lambda expression to create it’s
instance.  Since run() method takes no argument, our lambda expression also have no
argument.  Just like if-else blocks, we can avoid curly braces ({}) since we have a single
statement in the method body. For multiple statements, we would have to use curly braces
like any other methods.

Why do we need Lambda Expression?

1. Reduced Lines of Code - One of the clear benefit of using lambda expression is that the
    amount of code is reduced, we have already seen that how easily we can create instance
    of a functional interface using lambda expression rather than using anonymous class.
2. Sequential and Parallel Execution Support - Another benefit of using lambda expression
    is that we can benefit from the Stream API sequential and parallel operations support.

To explain this, let’s take a simple example where we need to write a method to test if a
number passed is prime number or not.

Traditionally we would write it’s code like below. The code is not fully optimized but
good for example purpose, so bear with me on this.

    //Traditional approach
    private static boolean isPrime(int number) {
        if(number < 2) return false;
        for(int i=2; i<number; i++){
            if(number % i == 0) return false;
        }
        return true;
    }

The problem with above code is that it's sequential in nature, if the number is very huge
then it will take significant amount of time. Another problem with code is that there are
so many exit points and it's not readable. Let's see how we can write the same method
using lambda expressions and stream API.

    //Declarative approach
    private static boolean isPrime(int number) {
        return number > 1
                && IntStream.range(2, number).noneMatch(
                        index -> number % index == 0);
    }

## Dependency Injection

- <https://dzone.com/articles/spring-di-patterns-the-good-the-bad-and-the-ugly>

## Type inference

What is Type Inference?
Is Type Inference available in older versions like Java 7 and Before 7
or it is available only in Java SE 8?

Type Inference means determining the Type by compiler at compile-time.

It is not new feature in Java SE 8. It is available in Java 7 and before Java 7 too.

**Before Java 7**:-
Let us explore Java arrays. Define a String of Array with values as shown below:

    String str[] = { "Java 7", "Java 8", "Java 9" };

Here we have assigned some String values at right side, but not defined it’s type. Java Compiler automatically infers it’s type and creates a String of Array.

**Java 7**:- Oracle Corporation has introduced “Diamond Operator” new feature in Java SE 7 to
avoid unnecessary Type definition in Generics.

    Map<String,List<Customer>> customerInfoByCity = new HashMap<>();

Here we have not defined Type information at right side, simply defined Java SE 7’s
Diamond Operator “”.

**Java SE 8**:- Oracle Corporation has enhanced this Type Inference concept a lot in Java
SE 8. We use this concept to define Lambda Expressions, Functions, Method References etc.

    ToIntBiFunction<Integer,Integer> add = (a,b) -> a + b;

Here Java Compiler observes the type definition available at left-side and determines the
type of Lambda Expression parameters a and b as Integers.

## Static import

The static import declaration is analogous to the normal import declaration. Where the
normal import declaration imports classes from packages, allowing them to be used without
package qualification, the static import declaration imports static members from classes,
allowing them to be used without class qualification.

So when should you use static import? Very sparingly! Only use it when you'd otherwise be
tempted to declare local copies of constants, or to abuse inheritance (the Constant
Interface Antipattern). In other words, use it when you require frequent access to static
members from one or two classes. If you overuse the static import feature, it can make
your program unreadable and unmaintainable, polluting its namespace with all the static
members you import. Readers of your code (including you, a few months after you wrote it)
will not know which class a static member comes from. Importing all of the static members
from a class can be particularly harmful to readability; if you need only one or two
members, import them individually. Used appropriately, static import can make your program
more readable, by removing the boilerplate of repetition of class names.

    import org.apache.commons.lang.StringUtils;

    if (StringUtils.isBlank(aString)) {

I can do this:

    import static org.apache.commons.lang.StringUtils.isBlank;

    if (isBlank(aString)) {
# WIKI Spring

## Beans - WTF ?

GENERAL:

“A Java Bean is a reusable software component that can be manipulated visually
in a builder tool.”

<https://stackoverflow.com/questions/3295496/what-is-a-javabean-exactly>

A JavaBean is just a standard

1. All properties private (use getters/setters)
2. A public no-argument constructor
3. Implements Serializable.

IN SPRING
The objects that form the backbone of your application and that are managed by the Spring
IoC container are called beans. A bean is an object that is instantiated, assembled, and
otherwise managed by a Spring IoC container. These beans are created with the
configuration metadata that you supply to the container, for example, in the form of XML
definitions.  More to learn about beans and scope from SpringSource:

When you create a bean definition what you are actually creating is a recipe for creating
actual instances of the class defined by that bean definition. The idea that a bean
definition is a recipe is important, because it means that, just like a class, you can
potentially have many object instances created from a single recipe.

You can control not only the various dependencies and configuration values that are to be
plugged into an object that is created from a particular bean definition, but also the
scope of the objects created from a particular bean definition. This approach is very
powerful and gives you the flexibility to choose the scope of the objects you create
through configuration instead of having to 'bake in' the scope of an object at the Java
class level. Beans can be defined to be deployed in one of a number of scopes

## Autowired

<https://stackoverflow.com/questions/19414734/understanding-spring-autowired-usage>
