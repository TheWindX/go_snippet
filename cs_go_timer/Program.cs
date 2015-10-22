using System;
using System.Collections.Generic;

class Program
{
    static int TimeSince(DateTime start)
    {
        var dur = (DateTime.Now - start);
        return (int)(dur.Ticks / 10000);
    }


    static IEnumerable<goroutinewrapper> NewTimer(chan endFlag, float t)
    {
        var start = DateTime.Now;
        var timeAdd = 0;
        Console.WriteLine("计时开始...0");
        for (;;)
        {
            if (TimeSince(start) > timeAdd)
            {
                timeAdd = TimeSince(start);
                if (TimeSince(start) % 500 == 0)
                    Console.WriteLine("计时中..." + TimeSince(start));
            }
            if (TimeSince(start) > t)
            {
                yield return goroutine.写(endFlag, TimeSince(start));
                break;
            }
            yield return null;
        }
    }

    static IEnumerable<goroutinewrapper> afterTimer(chan endFlag)
    {
        yield return goroutine.读(endFlag);
        Console.WriteLine("时间到达:{0}", endFlag.Val);
    }

    static void Main(string[] args)
    {
        var endFlag = chan.make();

        goroutine.go(NewTimer(endFlag, 3000));
        goroutine.go(afterTimer(endFlag));

        goroutine._for();
    }
}
