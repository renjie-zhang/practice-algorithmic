class Foo {

    public Foo() {

    }
     private AtomicInteger first = new AtomicInteger(0);
     private AtomicInteger second = new AtomicInteger(0);

    public void first(Runnable printFirst) throws InterruptedException {

        // printFirst.run() outputs "first". Do not change or remove this line.
        printFirst.run();
        first.incrementAndGet();
    }

    public void second(Runnable printSecond) throws InterruptedException {
        while(first.get() != 1){

        }
        // printSecond.run() outputs "second". Do not change or remove this line.
        printSecond.run();
        second.incrementAndGet();
    }

    public void third(Runnable printThird) throws InterruptedException {
        while(second.get() != 1){}
        // printThird.run() outputs "third". Do not change or remove this line.
        printThird.run();
    }
}