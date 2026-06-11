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

// ─── New feature tests ────────────────────────────────────────────────────────

// TestRouterPathParams verifies single-segment path parameter registration.
func TestRouterPathParams(t *testing.T) {
	input := `
dhoro app = router_banao();
app.ana("/users/:id", kaj(req, res) {
    dhoro id = req["params"]["id"];
    json_uttor(res, {"id": id});
});
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("path param route registration failed: %v", result.Inspect())
	}
	if result.Type() != object.MAP_OBJ {
		t.Fatalf("expected MAP (router), got %s", result.Type())
	}
}

// TestRouterPathParamsMultiple verifies multi-segment path parameters.
func TestRouterPathParamsMultiple(t *testing.T) {
	input := `
dhoro app = router_banao();
app.ana("/posts/:pid/comments/:cid", kaj(req, res) {
    dhoro pid = req["params"]["pid"];
    dhoro cid = req["params"]["cid"];
    json_uttor(res, {"pid": pid, "cid": cid});
});
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("multi-param route failed: %v", result.Inspect())
	}
}

// TestQueryParamsParsedAsMap verifies that routes can access req["query"]["key"].
func TestQueryParamsParsedAsMap(t *testing.T) {
	input := `
dhoro app = router_banao();
app.ana("/search", kaj(req, res) {
    dhoro term = req["query"]["q"];
    dhoro page = req["query"]["page"];
    json_uttor(res, {"term": term, "page": page});
});
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("query param route failed: %v", result.Inspect())
	}
}

// TestRouterMiddleware verifies app.majhe() accepts a 3-arg handler.
func TestRouterMiddleware(t *testing.T) {
	input := `
dhoro app = router_banao();
app.majhe(kaj(req, res, agorao) {
    dekho("middleware hit");
    agorao();
});
app.ana("/", kaj(req, res) {
    uttor(res, "OK");
});
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("middleware registration failed: %v", result.Inspect())
	}
}

// TestRouterMiddlewareChaining verifies multiple middleware layers.
func TestRouterMiddlewareChaining(t *testing.T) {
	input := `
dhoro app = router_banao();
app.majhe(kaj(req, res, agorao) { agorao(); });
app.majhe(kaj(req, res, agorao) { agorao(); });
app.majhe(kaj(req, res, agorao) { agorao(); });
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("chained middleware failed: %v", result.Inspect())
	}
}

// TestSubRouterMountingFixed verifies bebohar() correctly mounts a sub-router.
func TestSubRouterMountingFixed(t *testing.T) {
	input := `
dhoro api = router_banao();
api.ana("/users", kaj(req, res) { json_uttor(res, {"users": []}); });
api.pathano("/users", kaj(req, res) { json_uttor(res, {"created": sotti}, 201); });

dhoro app = router_banao();
app.bebohar("/api", api);
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("bebohar sub-router mounting failed: %v", result.Inspect())
	}
}

// TestCORSHelper verifies cors_chharpao() accepts default and custom options.
func TestCORSHelper(t *testing.T) {
	cases := []string{
		`dhoro app = router_banao(); cors_chharpao(app); app`,
		`dhoro app = router_banao(); cors_chharpao(app, {"origin": "https://example.com"}); app`,
		`dhoro app = router_banao(); cors_chharpao(app, {"origin": "*", "methods": "GET,POST"}); app`,
	}
	for _, input := range cases {
		result := testEval(input)
		if isError(result) {
			t.Fatalf("cors_chharpao failed: %v", result.Inspect())
		}
	}
}

// TestStaticFilesHelper verifies file_dao() registers without error.
func TestStaticFilesHelper(t *testing.T) {
	input := `
dhoro app = router_banao();
file_dao(app, "/public", ".");
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("file_dao failed: %v", result.Inspect())
	}
}

// TestGhurao verifies redirect sets 302 status.
func TestGhurao(t *testing.T) {
	input := `
dhoro res = {"status": 200, "body": "", "headers": {}};
ghurao(res, "/login");
res["status"]
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("ghurao failed: %v", result.Inspect())
	}
	if result.Type() != object.NUMBER_OBJ {
		t.Fatalf("expected NUMBER status, got %s", result.Type())
	}
	if result.(*object.Number).Value != 302 {
		t.Fatalf("expected status 302, got %v", result.(*object.Number).Value)
	}
}

// TestGhuraoCustomStatus verifies redirect with explicit 301.
func TestGhuraoCustomStatus(t *testing.T) {
	input := `
dhoro res = {"status": 200, "body": "", "headers": {}};
ghurao(res, "/new-page", 301);
res["status"]
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("ghurao 301 failed: %v", result.Inspect())
	}
	if result.(*object.Number).Value != 301 {
		t.Fatalf("expected status 301, got %v", result.(*object.Number).Value)
	}
}

// TestKukiRakho verifies kuki_rakho() sets a Set-Cookie header.
func TestKukiRakho(t *testing.T) {
	input := `
dhoro res = {"status": 200, "body": "", "headers": {}};
kuki_rakho(res, "token", "abc123");
res["headers"]["Set-Cookie"]
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("kuki_rakho failed: %v", result.Inspect())
	}
	if result.Type() != object.STRING_OBJ {
		t.Fatalf("expected STRING Set-Cookie, got %s", result.Type())
	}
	if result.(*object.String).Value == "" {
		t.Fatal("Set-Cookie header should not be empty")
	}
}

// TestKukiRakhoWithOptions verifies kuki_rakho() with all cookie attributes.
func TestKukiRakhoWithOptions(t *testing.T) {
	input := `
dhoro res = {"status": 200, "body": "", "headers": {}};
kuki_rakho(res, "session", "xyz789", {"httpOnly": sotti, "maxAge": 3600, "path": "/app", "sameSite": "Lax"});
res["headers"]["Set-Cookie"]
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("kuki_rakho with options failed: %v", result.Inspect())
	}
	if result.Type() != object.STRING_OBJ {
		t.Fatalf("expected STRING Set-Cookie, got %s", result.Type())
	}
}

// TestGotiShima verifies goti_shima() configures the rate limiter without error.
func TestGotiShima(t *testing.T) {
	input := `
dhoro app = router_banao();
goti_shima(app, 100, 60);
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("goti_shima failed: %v", result.Inspect())
	}
}

// TestSankochonChalu verifies sankochon_chalu() enables gzip without error.
func TestSankochonChalu(t *testing.T) {
	input := `
dhoro app = router_banao();
sankochon_chalu(app);
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("sankochon_chalu failed: %v", result.Inspect())
	}
}

// TestSomoyShima verifies somoy_shima() sets the timeout without error.
func TestSomoyShima(t *testing.T) {
	input := `
dhoro app = router_banao();
somoy_shima(app, 30);
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("somoy_shima failed: %v", result.Inspect())
	}
}

// TestAkaarShima verifies akaar_shima() sets the body limit without error.
func TestAkaarShima(t *testing.T) {
	input := `
dhoro app = router_banao();
akaar_shima(app, 1048576);
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("akaar_shima failed: %v", result.Inspect())
	}
}

// TestBhulSambhalo verifies bhul_sambhalo() registers an error handler without error.
func TestBhulSambhalo(t *testing.T) {
	input := `
dhoro app = router_banao();
bhul_sambhalo(app, kaj(err, req, res) {
    json_uttor(res, {"error": "something went wrong"}, 500);
});
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("bhul_sambhalo failed: %v", result.Inspect())
	}
}

// TestLogChalu verifies log_chalu() enables logging without error.
func TestLogChalu(t *testing.T) {
	input := `
dhoro app = router_banao();
log_chalu(app);
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("log_chalu failed: %v", result.Inspect())
	}
}

// TestFullProductionStack verifies all middleware helpers chain without error.
func TestFullProductionStack(t *testing.T) {
	input := `
dhoro app = router_banao();
cors_chharpao(app);
log_chalu(app);
sankochon_chalu(app);
somoy_shima(app, 30);
akaar_shima(app, 1048576);
goti_shima(app, 100, 60);
app.majhe(kaj(req, res, agorao) { agorao(); });
file_dao(app, "/public", ".");
app.ana("/users/:id", kaj(req, res) {
    json_uttor(res, {"id": req["params"]["id"]});
});
app.pathano("/users", kaj(req, res) {
    dhoro body = req["json"];
    json_uttor(res, {"created": sotti}, 201);
});
app.ana("/search", kaj(req, res) {
    json_uttor(res, {"q": req["query"]["q"]});
});
bhul_sambhalo(app, kaj(err, req, res) {
    json_uttor(res, {"error": "internal"}, 500);
});
app
`
	result := testEval(input)
	if isError(result) {
		t.Fatalf("full production stack failed: %v", result.Inspect())
	}
	if result.Type() != object.MAP_OBJ {
		t.Fatalf("expected MAP (router), got %s", result.Type())
	}
}
