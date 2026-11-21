namespace DddTemplate.Tests.Unit.Domain;

using DddTemplate.Domain;
using Xunit;

[Trait("Category", "Unit")]
public class ValueObjectTests
{
    [Fact]
    public void EqualOperator_HandlesNullAndEqualValues()
    {
        var value = new TestValueObject("one");

        Assert.False(TestValueObject.Same(null, value));
        Assert.False(TestValueObject.Same(value, null));
        Assert.True(TestValueObject.Same(null, null));
        Assert.True(TestValueObject.Same(value, new TestValueObject("one")));
        Assert.False(TestValueObject.Same(value, new TestValueObject("two")));
    }

    [Fact]
    public void NotEqualOperator_InvertsEquality()
    {
        var value = new TestValueObject("one");

        Assert.True(TestValueObject.Different(null, value));
        Assert.True(TestValueObject.Different(value, null));
        Assert.False(TestValueObject.Different(null, null));
        Assert.False(TestValueObject.Different(value, new TestValueObject("one")));
        Assert.True(TestValueObject.Different(value, new TestValueObject("two")));
    }

    private sealed class TestValueObject(string value) : ValueObject
    {
        public static bool Same(ValueObject? left, ValueObject? right) =>
            EqualOperator(left!, right!);

        public static bool Different(ValueObject? left, ValueObject? right) =>
            NotEqualOperator(left!, right!);

        public override bool Equals(object? obj) =>
            obj is TestValueObject other && value == other.Value;

        private string Value => value;

        public override int GetHashCode() => value.GetHashCode();
    }
}
