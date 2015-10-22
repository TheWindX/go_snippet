using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;


class goroutinewrapper
{
    public bool readOrWrite = false;
    public bool ready()
    {
        if (readOrWrite) return inner.canRead();
        else return inner.canWrite();
    }
    public chan inner;
}

class chan
{

    public object _value = null;
    public object Val
    {
        get
        {
            var r = _value;
            _value = null;
            return r;
        }
        set
        {
            _value = value;
        }
    }

    public bool canRead()
    {
        if (_value == null) return false;
        return true;
    }

    public bool canWrite()
    {
        if (_value == null) return true;
        return false;
    }

    public static chan make()
    {
        var r = new chan();
        return r;
    }
}


class goroutine
{
    public static goroutinewrapper 读(chan o)
    {
        var r = new goroutinewrapper();
        r.inner = o;
        r.readOrWrite = true;
        return r;
    }

    public static goroutinewrapper 写(chan o, object obj)
    {
        var r = new goroutinewrapper();
        r.inner = o;
        r.inner.Val = obj;
        r.readOrWrite = false;
        return r;
    }

    public static void go(IEnumerable<goroutinewrapper> routine)
    {
        var iter = routine.GetEnumerator();
        _waiting.Add(iter);
    }

    public static List<IEnumerator<goroutinewrapper>> _waiting = new List<IEnumerator<goroutinewrapper>>();

    public static bool moveNext()
    {
        bool changed = false;
        List<IEnumerator<goroutinewrapper>> _todels = new List<IEnumerator<goroutinewrapper>>();

        foreach (var item in _waiting)
        {
            var c = item.Current;
            if (c == null)
            {
                if (!item.MoveNext())
                {
                    _todels.Add(item);
                }
                continue;
            }
            if (c.ready())
            {
                if (!item.MoveNext())
                {
                    _todels.Add(item);
                }
                changed = true;
            }
        }

        foreach (var item in _todels)
        {
            _waiting.Remove(item);
        }
        return changed;
    }

    public static void _for()
    {
        for (; ; )
        {
            while (moveNext()) { }
            if (_waiting.Count == 0)
                return;
        }   
    }
}