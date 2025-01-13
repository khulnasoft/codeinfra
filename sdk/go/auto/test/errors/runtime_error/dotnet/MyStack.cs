using Codeinfra;
using System;

class MyStack : Stack
{
    public MyStack()
    {
       throw new ApplicationException("thrown for testing");
    }
}
