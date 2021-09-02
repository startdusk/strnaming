package strnaming

import "testing"

func TestToCamel(t *testing.T) {
	testToCamel(t)
}

func BenchmarkToCamel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testToCamel(b)
	}
}

func testToCamel(tb testing.TB) {
	cases := []struct {
		test       string
		expect     string
		lowerFirst bool
		split      byte
	}{
		{
			test:   "_abc_def",
			expect: "AbcDef",
		},
		{
			test:   "           ***_abc_def",
			expect: "AbcDef",
		},
		{
			test:   "   ***_abc_DB_",
			expect: "AbcDB",
		},
		{
			test:   "DNV.abc",
			expect: "DNVAbc",
			split:  '.',
		},
		{
			test:   "test_case",
			expect: "TestCase",
			split:  '_',
		},
		{
			test:   "many2many",
			expect: "Many2Many",
		},
		{
			test:   "department_id",
			expect: "DepartmentId",
		},
		{
			test:       "department_id",
			expect:     "departmentId",
			lowerFirst: true,
		},
		{
			test:       "_Department_id",
			expect:     "departmentId",
			lowerFirst: true,
		},
		{
			test:   "_Department_id",
			expect: "DepartmentId",
		},
		{
			test:       "__prod__",
			expect:     "prod",
			lowerFirst: true,
		},
		{
			test:   "",
			expect: "",
		},
		{
			test:       "Golang",
			expect:     "golang",
			lowerFirst: true,
		},
		{
			test:   "__foo_ Fzz**",
			expect: "FooFzz",
		},
		{
			test:       "__foo_ Fzz**",
			expect:     "fooFzz",
			lowerFirst: true,
		},
	}

	for _, cc := range cases {
		actual := ToCamel(cc.test).WithLowerFirst(cc.lowerFirst).WithSplit(cc.split).String()
		if actual != cc.expect {
			tb.Errorf("expect camel case %s, but got %s\n", cc.expect, actual)
		}
	}
}
