using System;
using Grpc.Core;
using Demo.CalcPb;
using System.Threading.Tasks;

namespace cs_client
{
    class Program
    {
        static async Task Main(string[] args)
        {
            Channel channel = new Channel("127.0.0.1:50051", ChannelCredentials.Insecure);
            var client = new CalculatorService.CalculatorServiceClient(channel);

            var result = await client.AddAsync(new AddRequest
            {
                No1 = 1,
                No2 = 2,
            });

            var result1 = await client.AddAsync(new AddRequest
            {
                No1 = 2,
                No2 = 5,
            });

            var result2 = await client.AddAsync(new AddRequest
            {
                No1 = 5,
                No2 = 10,
            });

            Console.WriteLine(result);
            Console.WriteLine(result1);
            Console.WriteLine(result2);

            await channel.ShutdownAsync();
            Console.WriteLine("Press any key to exit...");
            Console.ReadKey();
        }
    }
}
