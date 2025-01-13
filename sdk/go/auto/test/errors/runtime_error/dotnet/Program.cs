using System.Threading.Tasks;
using Codeinfra;

class Program
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
