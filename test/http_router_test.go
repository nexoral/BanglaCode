package test

import (
	"BanglaCode/src/object"
	"testing"
)

// Test basic router creation
func TestRouterCreation(t *testing.T) {
	input := `
dhoro router = router_banao();
dekho(dhoron(router));
`

	result := testEval(input)
	if result == nil {
		t.Fatal("Expected router object, got nil")
	}
}

// Test router GET method
func TestRouterGetMethod(t *testing.T) {
	input := `
dhoro router = router_banao();
router.ana("/test", kaj(req, res) {
    uttor(res, "test response");
});
dekho("Router created with GET route");
`

	result := testEval(input)
	if result == nil {
		t.Fatal("Expected successful route registration")
	}
}

// Test router POST method
func TestRouterPostMethod(t *testing.T) {
	input := `
dhoro router = router_banao();
router.pathano("/submit", kaj(req, res) {
    json_uttor(res, {"status": "success"});
});
dekho("Router created with POST route");
`

	result := testEval(input)
	if result == nil {
		t.Fatal("Expected successful route registration")
	}
}

// Test router PUT method
func TestRouterPutMethod(t *testing.T) {
	input := `
dhoro router = router_banao();
router.bodlano("/update", kaj(req, res) {
    uttor(res, "updated", 200);
});
dekho("Router created with PUT route");
`

	result := testEval(input)
	if result == nil {
		t.Fatal("Expected successful route registration")
	}
}

// Test router DELETE method
func TestRouterDeleteMethod(t *testing.T) {
	input := `
dhoro router = router_banao();
router.mujhe_felo("/remove", kaj(req, res) {
    uttor(res, "deleted", 204);
});
dekho("Router created with DELETE route");
`

	result := testEval(input)
	if result == nil {
		t.Fatal("Expected successful route registration")
	}
}

// Test router PATCH method
func TestRouterPatchMethod(t *testing.T) {
	input := `
dhoro router = router_banao();
router.songshodhon("/modify", kaj(req, res) {
    uttor(res, "modified", 200);
});
dekho("Router created with PATCH route");
`

	result := testEval(input)
	if result == nil {
		t.Fatal("Expected successful route registration")
	}
}

// Test router method chaining
func TestRouterMethodChaining(t *testing.T) {
	input := `
dhoro router = router_banao();

router
    .ana("/", kaj(req, res) { uttor(res, "home"); })
    .ana("/about", kaj(req, res) { uttor(res, "about"); })
    .pathano("/submit", kaj(req, res) { uttor(res, "submitted"); });

dekho("Router created with chained methods");
`

	result := testEval(input)
	if result == nil {
		t.Fatal("Expected successful method chaining")
	}
}

// Test multiple routers
func TestMultipleRouters(t *testing.T) {
	input := `
dhoro authRouter = router_banao();
authRouter.pathano("/login", kaj(req, res) {
    json_uttor(res, {"token": "abc123"});
});

dhoro userRouter = router_banao();
userRouter.ana("/", kaj(req, res) {
    json_uttor(res, {"users": []});
});

dekho("Multiple routers created successfully");
`

	result := testEval(input)
	if result == nil {
		t.Fatal("Expected successful creation of multiple routers")
	}
}

// Test router with complex handler
func TestRouterComplexHandler(t *testing.T) {
	input := `
dhoro router = router_banao();

router.pathano("/api/data", kaj(req, res) {
    dhoro method = req["method"];
    dhoro path = req["path"];
    dhoro body = req["body"];

    dhoro response = {
        "received": {
            "method": method,
            "path": path,
            "body": body
        },
        "processed": sotti
    };

    json_uttor(res, response, 201);
});

dekho("Router with complex handler created");
`

	result := testEval(input)
	if result == nil {
		t.Fatal("Expected successful complex handler registration")
	}
}

// Test router error handling
func TestRouterErrorHandling(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name: "Missing path argument",
			input: `
dhoro router = router_banao();
router.ana();
`,
			wantErr: true,
		},
		{
			name: "Invalid path type",
			input: `
dhoro router = router_banao();
router.ana(123, kaj(req, res) {});
`,
			wantErr: true,
		},
		{
			name: "Invalid handler type",
			input: `
dhoro router = router_banao();
router.ana("/test", "not a function");
`,
			wantErr: true,
		},
		{
			name: "Too many arguments",
			input: `
dhoro router = router_banao();
router.ana("/test", kaj(req, res) {}, "extra");
`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := testEval(tt.input)
			if tt.wantErr && !isError(result) {
				t.Errorf("Expected error for %s, got %v", tt.name, result)
			}
		})
	}
}

func isError(result object.Object) bool {
	if result == nil {
		return false
	}
	return result.Type() == object.ERROR_OBJ
}

// Benchmark router creation
func BenchmarkRouterCreation(b *testing.B) {
	input := `dhoro router = router_banao();`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testEval(input)
	}
}

// Benchmark route registration
func BenchmarkRouteRegistration(b *testing.B) {
	input := `
dhoro router = router_banao();
router.ana("/test", kaj(req, res) {
    uttor(res, "test");
});
`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testEval(input)
	}
}

// Benchmark multiple routes
func BenchmarkMultipleRoutes(b *testing.B) {
	input := `
dhoro router = router_banao();
router.ana("/", kaj(req, res) { uttor(res, "home"); });
router.ana("/about", kaj(req, res) { uttor(res, "about"); });
router.pathano("/submit", kaj(req, res) { uttor(res, "ok"); });
router.bodlano("/update", kaj(req, res) { uttor(res, "updated"); });
router.mujhe_felo("/delete", kaj(req, res) { uttor(res, "deleted"); });
`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testEval(input)
	}
}
