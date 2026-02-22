export const metadata = {
  title: 'URL Parsing API - BanglaCode',
  description: 'Learn how to parse, manipulate, and construct URLs in BanglaCode with comprehensive URL and query string handling.',
};

export default function URLParsingDoc() {
  return (
    <div className="max-w-4xl mx-auto px-6 py-10">
      <h1 className="text-4xl font-bold mb-6">URL Parsing API</h1>
      
      <p className="text-lg text-gray-700 dark:text-gray-300 mb-8">
        The URL Parsing API in BanglaCode provides powerful tools for parsing, manipulating, and constructing URLs. Whether you're building APIs, handling query parameters, or working with web services, these functions make URL operations simple and reliable.
      </p>

      {/* Quick Start */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-4">Quick Start</h2>
        <div className="bg-gray-100 dark:bg-gray-800 rounded-lg p-6">
          <pre className="text-sm overflow-x-auto">
            <code className="language-banglacode">
{`// Parse a URL
dhoro url = url_parse("https://api.example.com:8080/users?role=admin&page=2#results");

dekho("Protocol:", url.Protocol);    // "https:"
dekho("Hostname:", url.Hostname);    // "api.example.com"
dekho("Port:", url.Port);            // "8080"
dekho("Pathname:", url.Pathname);    // "/users"
dekho("Search:", url.Search);        // "?role=admin&page=2"
dekho("Hash:", url.Hash);            // "#results"

// Work with query parameters
dhoro params = url_query_params(url.Search);
dekho("Role:", url_query_get(params, "role"));     // "admin"
dekho("Page:", url_query_get(params, "page"));     // "2"`}
            </code>
          </pre>
        </div>
      </section>

      {/* Core Concepts */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-4">Core Concepts</h2>
        
        <div className="space-y-6">
          <div className="border-l-4 border-blue-500 pl-6">
            <h3 className="text-xl font-semibold mb-2">URL Object (ইউআরএল অবজেক্ট)</h3>
            <p className="text-gray-700 dark:text-gray-300">
              A URL object represents a parsed URL with all its components: protocol, hostname, port, pathname, query string, and hash. It provides structured access to each part of a URL for easy manipulation.
            </p>
          </div>

          <div className="border-l-4 border-green-500 pl-6">
            <h3 className="text-xl font-semibold mb-2">URLSearchParams (কোয়েরি প্যারামিটার)</h3>
            <p className="text-gray-700 dark:text-gray-300">
              URLSearchParams provides a convenient interface for working with URL query strings. It supports getting, setting, appending, and deleting parameters, making query string manipulation straightforward and error-free.
            </p>
          </div>

          <div className="border-l-4 border-purple-500 pl-6">
            <h3 className="text-xl font-semibold mb-2">Query String Manipulation</h3>
            <p className="text-gray-700 dark:text-gray-300">
              Query strings encode key-value pairs in URLs (e.g., ?name=value&key=data). The API handles encoding/decoding automatically, supports multiple values per key, and provides methods to iterate over all parameters.
            </p>
          </div>

          <div className="border-l-4 border-orange-500 pl-6">
            <h3 className="text-xl font-semibold mb-2">URL Encoding</h3>
            <p className="text-gray-700 dark:text-gray-300">
              Special characters in URLs must be encoded. The API automatically handles URL encoding when setting parameters, ensuring your URLs are always valid and safe to use in HTTP requests.
            </p>
          </div>
        </div>
      </section>

      {/* API Reference */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6">API Reference</h2>

        <div className="space-y-8">
          {/* url_parse */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-blue-600 dark:text-blue-400">
              url_parse(urlString)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Parses a URL string and returns a URL object with all its components.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>urlString</code> (string): The URL to parse</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">URL object with properties:</p>
              <ul className="list-disc list-inside text-sm space-y-1 ml-4 mt-2">
                <li><code>Href</code>: Full URL string</li>
                <li><code>Protocol</code>: Protocol scheme (e.g., "https:")</li>
                <li><code>Username</code>: Username for authentication</li>
                <li><code>Password</code>: Password for authentication</li>
                <li><code>Hostname</code>: Domain name or IP address</li>
                <li><code>Port</code>: Port number (if specified)</li>
                <li><code>Host</code>: Hostname + port (if present)</li>
                <li><code>Pathname</code>: Path portion of the URL</li>
                <li><code>Search</code>: Query string (including "?")</li>
                <li><code>Hash</code>: Fragment identifier (including "#")</li>
                <li><code>Origin</code>: Protocol + hostname + port</li>
              </ul>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro url = url_parse("https://user:pass@api.example.com:8080/v1/users?active=true#top");

dekho(url.Href);      // "https://user:pass@api.example.com:8080/v1/users?active=true#top"
dekho(url.Protocol);  // "https:"
dekho(url.Username);  // "user"
dekho(url.Password);  // "pass"
dekho(url.Hostname);  // "api.example.com"
dekho(url.Port);      // "8080"
dekho(url.Host);      // "api.example.com:8080"
dekho(url.Pathname);  // "/v1/users"
dekho(url.Search);    // "?active=true"
dekho(url.Hash);      // "#top"
dekho(url.Origin);    // "https://api.example.com:8080"`}
                </code>
              </pre>
            </div>
          </div>

          {/* url_query_params */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-green-600 dark:text-green-400">
              url_query_params(queryStringOrURL)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Creates a URLSearchParams object from a query string or URL object.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>queryStringOrURL</code> (string or URL object): Query string (with or without "?") or URL object</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">URLSearchParams object for manipulating query parameters</p>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// From query string
dhoro params1 = url_query_params("?name=John&age=30");
dhoro params2 = url_query_params("name=John&age=30");  // Works without "?"

// From URL object
dhoro url = url_parse("https://api.com/users?role=admin");
dhoro params3 = url_query_params(url.Search);

// Empty params
dhoro params4 = url_query_params("");  // Create empty params`}
                </code>
              </pre>
            </div>
          </div>

          {/* url_query_get */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-purple-600 dark:text-purple-400">
              url_query_get(params, key)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Gets the value of a query parameter. Returns the first value if multiple exist.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>params</code> (URLSearchParams): The params object</li>
                <li><code>key</code> (string): The parameter name</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">String value of the parameter, or null if not found</p>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro params = url_query_params("?name=John&age=30&tag=js&tag=go");

dhoro name = url_query_get(params, "name");    // "John"
dhoro age = url_query_get(params, "age");      // "30"
dhoro tag = url_query_get(params, "tag");      // "js" (first value)
dhoro missing = url_query_get(params, "city"); // null

jodi (name != mittha) {
  dekho("Name is:", name);
} nahole {
  dekho("Name not found");
}`}
                </code>
              </pre>
            </div>
          </div>

          {/* url_query_set */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-orange-600 dark:text-orange-400">
              url_query_set(params, key, value)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Sets a query parameter value. Replaces all existing values for that key.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>params</code> (URLSearchParams): The params object</li>
                <li><code>key</code> (string): The parameter name</li>
                <li><code>value</code> (string): The value to set</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">None (modifies params in place)</p>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro params = url_query_params("?name=John&age=30");

// Set new parameter
url_query_set(params, "city", "Boston");
dekho(url_query_toString(params));  // "name=John&age=30&city=Boston"

// Replace existing parameter
url_query_set(params, "age", "31");
dekho(url_query_toString(params));  // "name=John&age=31&city=Boston"

// Set replaces all values if multiple exist
url_query_set(params, "tag", "first");
url_query_append(params, "tag", "second");
url_query_set(params, "tag", "only");  // Removes "first" and "second"
dekho(url_query_get(params, "tag"));   // "only"`}
                </code>
              </pre>
            </div>
          </div>

          {/* url_query_append */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-pink-600 dark:text-pink-400">
              url_query_append(params, key, value)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Appends a query parameter value. Allows multiple values for the same key.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>params</code> (URLSearchParams): The params object</li>
                <li><code>key</code> (string): The parameter name</li>
                <li><code>value</code> (string): The value to append</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">None (modifies params in place)</p>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro params = url_query_params("?name=John");

// Append multiple values for same key
url_query_append(params, "tag", "javascript");
url_query_append(params, "tag", "golang");
url_query_append(params, "tag", "python");

dekho(url_query_toString(params));
// "name=John&tag=javascript&tag=golang&tag=python"

// Get only returns first value
dekho(url_query_get(params, "tag"));  // "javascript"

// Get all values
dhoro allTags = url_query_values(params);  // ["John", "javascript", "golang", "python"]`}
                </code>
              </pre>
            </div>
          </div>

          {/* url_query_delete */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-red-600 dark:text-red-400">
              url_query_delete(params, key)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Removes a query parameter and all its values.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>params</code> (URLSearchParams): The params object</li>
                <li><code>key</code> (string): The parameter name to remove</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">None (modifies params in place)</p>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro params = url_query_params("?name=John&age=30&city=Boston");

dekho(url_query_toString(params));
// "name=John&age=30&city=Boston"

// Delete a parameter
url_query_delete(params, "age");
dekho(url_query_toString(params));
// "name=John&city=Boston"

// Delete removes all values if multiple exist
url_query_append(params, "tag", "js");
url_query_append(params, "tag", "go");
url_query_delete(params, "tag");  // Removes both values
dekho(url_query_has(params, "tag"));  // mittha`}
                </code>
              </pre>
            </div>
          </div>

          {/* url_query_has */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-indigo-600 dark:text-indigo-400">
              url_query_has(params, key)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Checks if a query parameter exists.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>params</code> (URLSearchParams): The params object</li>
                <li><code>key</code> (string): The parameter name to check</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">Boolean: shotti (true) if parameter exists, mittha (false) otherwise</p>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro params = url_query_params("?name=John&age=30");

jodi (url_query_has(params, "name")) {
  dekho("Name exists:", url_query_get(params, "name"));
}

jodi (url_query_has(params, "email")) {
  dekho("Email exists");
} nahole {
  dekho("Email not found");  // This will print
}

// Use for conditional logic
jodi (!url_query_has(params, "page")) {
  url_query_set(params, "page", "1");  // Set default page
}`}
                </code>
              </pre>
            </div>
          </div>

          {/* url_query_keys */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-teal-600 dark:text-teal-400">
              url_query_keys(params)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Returns an array of all query parameter keys.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>params</code> (URLSearchParams): The params object</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">Array of strings containing all parameter names (includes duplicates)</p>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro params = url_query_params("?name=John&age=30&city=Boston");

dhoro keys = url_query_keys(params);
dekho(keys);  // ["name", "age", "city"]

// Iterate over keys
ghuriye (dhoro i = 0; i < dorghyo(keys); i = i + 1) {
  dhoro key = keys[i];
  dhoro value = url_query_get(params, key);
  dekho(key, "=", value);
}

// Multiple values for same key
url_query_append(params, "tag", "js");
url_query_append(params, "tag", "go");
dhoro allKeys = url_query_keys(params);
// ["name", "age", "city", "tag", "tag"] - duplicates included`}
                </code>
              </pre>
            </div>
          </div>

          {/* url_query_values */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-cyan-600 dark:text-cyan-400">
              url_query_values(params)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Returns an array of all query parameter values.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>params</code> (URLSearchParams): The params object</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">Array of strings containing all parameter values</p>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro params = url_query_params("?name=John&age=30&city=Boston");

dhoro values = url_query_values(params);
dekho(values);  // ["John", "30", "Boston"]

// Get all values including duplicates
url_query_append(params, "tag", "javascript");
url_query_append(params, "tag", "golang");
dhoro allValues = url_query_values(params);
// ["John", "30", "Boston", "javascript", "golang"]

// Iterate over values
ghuriye (dhoro i = 0; i < dorghyo(values); i = i + 1) {
  dekho("Value", i, ":", values[i]);
}`}
                </code>
              </pre>
            </div>
          </div>

          {/* url_query_toString */}
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-2xl font-semibold mb-3 text-amber-600 dark:text-amber-400">
              url_query_toString(params)
            </h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Converts URLSearchParams back to a query string.
            </p>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Parameters:</p>
              <ul className="list-disc list-inside text-sm space-y-1">
                <li><code>params</code> (URLSearchParams): The params object</li>
              </ul>
            </div>
            <div className="bg-gray-100 dark:bg-gray-900 rounded p-4 mb-4">
              <p className="text-sm font-semibold mb-2">Returns:</p>
              <p className="text-sm">String: URL-encoded query string (without leading "?")</p>
            </div>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`dhoro params = url_query_params("");

url_query_set(params, "name", "John Doe");
url_query_set(params, "age", "30");
url_query_set(params, "city", "New York");

dhoro queryString = url_query_toString(params);
dekho(queryString);  // "name=John+Doe&age=30&city=New+York"

// Build complete URL
dhoro baseUrl = "https://api.example.com/users";
dhoro fullUrl = baseUrl + "?" + queryString;
dekho(fullUrl);
// "https://api.example.com/users?name=John+Doe&age=30&city=New+York"

// Special characters are automatically encoded
url_query_set(params, "search", "hello world & more");
dekho(url_query_toString(params));
// "name=John+Doe&age=30&city=New+York&search=hello+world+%26+more"`}
                </code>
              </pre>
            </div>
          </div>
        </div>
      </section>

      {/* Real-World Examples */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6">Real-World Examples</h2>

        <div className="space-y-8">
          {/* Example 1: API Request Building */}
          <div className="bg-gradient-to-r from-blue-50 to-cyan-50 dark:from-blue-900/20 dark:to-cyan-900/20 rounded-lg p-6">
            <h3 className="text-2xl font-semibold mb-4">Example 1: API Request Building with Query Parameters</h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Build API requests dynamically with filters, pagination, and sorting:
            </p>
            <div className="bg-white dark:bg-gray-800 rounded-lg p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Build API URL with dynamic filters
kaj buildAPIRequest(baseUrl, filters) {
  dhoro params = url_query_params("");
  
  // Add filters if provided
  jodi (url_query_has(filters, "status")) {
    url_query_set(params, "status", url_query_get(filters, "status"));
  }
  
  jodi (url_query_has(filters, "role")) {
    url_query_set(params, "role", url_query_get(filters, "role"));
  }
  
  // Add pagination
  dhoro page = url_query_has(filters, "page") ? 
    url_query_get(filters, "page") : "1";
  dhoro limit = url_query_has(filters, "limit") ? 
    url_query_get(filters, "limit") : "20";
  
  url_query_set(params, "page", page);
  url_query_set(params, "limit", limit);
  
  // Add sorting
  jodi (url_query_has(filters, "sortBy")) {
    url_query_set(params, "sortBy", url_query_get(filters, "sortBy"));
    url_query_set(params, "order", 
      url_query_has(filters, "order") ? 
      url_query_get(filters, "order") : "asc"
    );
  }
  
  dhoro queryString = url_query_toString(params);
  ferao baseUrl + "?" + queryString;
}

// Usage
dhoro userFilters = url_query_params("?status=active&role=admin&page=2&sortBy=name");
dhoro apiUrl = buildAPIRequest("https://api.example.com/users", userFilters);
dekho(apiUrl);
// "https://api.example.com/users?status=active&role=admin&page=2&limit=20&sortBy=name&order=asc"

// Make API request
dhoro response = http_get(apiUrl);
dekho("Found", dorghyo(response.users), "users");`}
                </code>
              </pre>
            </div>
          </div>

          {/* Example 2: GitHub API URL Parsing */}
          <div className="bg-gradient-to-r from-green-50 to-emerald-50 dark:from-green-900/20 dark:to-emerald-900/20 rounded-lg p-6">
            <h3 className="text-2xl font-semibold mb-4">Example 2: GitHub API URL Parsing and Manipulation</h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Parse and modify GitHub API URLs for repository operations:
            </p>
            <div className="bg-white dark:bg-gray-800 rounded-lg p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Parse GitHub repository URL
kaj analyzeGitHubURL(githubUrl) {
  dhoro url = url_parse(githubUrl);
  
  // Extract repository information
  dhoro pathParts = bibhajan(url.Pathname, "/");
  dhoro owner = pathParts[1];
  dhoro repo = pathParts[2];
  
  dekho("Repository:", owner + "/" + repo);
  dekho("API Base:", url.Origin);
  
  // Parse query parameters
  dhoro params = url_query_params(url.Search);
  
  jodi (url_query_has(params, "page")) {
    dekho("Page:", url_query_get(params, "page"));
  }
  
  jodi (url_query_has(params, "per_page")) {
    dekho("Per Page:", url_query_get(params, "per_page"));
  }
  
  // Build issues API URL
  dhoro issuesParams = url_query_params("");
  url_query_set(issuesParams, "state", "open");
  url_query_set(issuesParams, "labels", "bug");
  url_query_set(issuesParams, "sort", "created");
  url_query_set(issuesParams, "direction", "desc");
  
  dhoro issuesUrl = url.Origin + "/repos/" + owner + "/" + repo + 
    "/issues?" + url_query_toString(issuesParams);
  
  dekho("Issues URL:", issuesUrl);
  ferao issuesUrl;
}

// Usage
dhoro repoUrl = "https://api.github.com/repos/golang/go?page=1&per_page=30";
dhoro issuesEndpoint = analyzeGitHubURL(repoUrl);
// Repository: golang/go
// API Base: https://api.github.com
// Page: 1
// Per Page: 30
// Issues URL: https://api.github.com/repos/golang/go/issues?state=open&labels=bug&sort=created&direction=desc

// Fetch issues
dhoro response = http_get(issuesEndpoint);
dekho("Found", dorghyo(response), "open bug issues");`}
                </code>
              </pre>
            </div>
          </div>

          {/* Example 3: Search/Filter URL Construction */}
          <div className="bg-gradient-to-r from-purple-50 to-pink-50 dark:from-purple-900/20 dark:to-pink-900/20 rounded-lg p-6">
            <h3 className="text-2xl font-semibold mb-4">Example 3: Search and Filter URL Construction</h3>
            <p className="text-gray-700 dark:text-gray-300 mb-4">
              Build complex search URLs with multiple filters and tags:
            </p>
            <div className="bg-white dark:bg-gray-800 rounded-lg p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Build search URL with multiple filters
kaj buildSearchURL(searchTerm, tags, options) {
  dhoro baseUrl = "https://example.com/search";
  dhoro params = url_query_params("");
  
  // Add search term (encoded automatically)
  jodi (searchTerm != "" && searchTerm != mittha) {
    url_query_set(params, "q", searchTerm);
  }
  
  // Add multiple tags
  jodi (tags != mittha && dorghyo(tags) > 0) {
    ghuriye (dhoro i = 0; i < dorghyo(tags); i = i + 1) {
      url_query_append(params, "tag", tags[i]);
    }
  }
  
  // Add optional filters
  jodi (options != mittha) {
    jodi (url_query_has(options, "category")) {
      url_query_set(params, "category", url_query_get(options, "category"));
    }
    
    jodi (url_query_has(options, "minPrice")) {
      url_query_set(params, "minPrice", url_query_get(options, "minPrice"));
    }
    
    jodi (url_query_has(options, "maxPrice")) {
      url_query_set(params, "maxPrice", url_query_get(options, "maxPrice"));
    }
    
    jodi (url_query_has(options, "sortBy")) {
      url_query_set(params, "sortBy", url_query_get(options, "sortBy"));
    }
  }
  
  // Build final URL
  dhoro queryString = url_query_toString(params);
  ferao queryString != "" ? baseUrl + "?" + queryString : baseUrl;
}

// Usage 1: Simple search
dhoro url1 = buildSearchURL("laptop computers", mittha, mittha);
dekho(url1);
// "https://example.com/search?q=laptop+computers"

// Usage 2: Search with tags
dhoro tags = ["electronics", "computers", "portable"];
dhoro url2 = buildSearchURL("gaming laptop", tags, mittha);
dekho(url2);
// "https://example.com/search?q=gaming+laptop&tag=electronics&tag=computers&tag=portable"

// Usage 3: Full featured search
dhoro filters = url_query_params("?category=electronics&minPrice=500&maxPrice=2000&sortBy=price");
dhoro url3 = buildSearchURL("laptop", tags, filters);
dekho(url3);
// "https://example.com/search?q=laptop&tag=electronics&tag=computers&tag=portable&category=electronics&minPrice=500&maxPrice=2000&sortBy=price"

// Parse existing search URL to modify
dhoro existingUrl = url_parse(url3);
dhoro existingParams = url_query_params(existingUrl.Search);

// Modify filters
url_query_set(existingParams, "maxPrice", "1500");
url_query_delete(existingParams, "sortBy");
url_query_set(existingParams, "sortBy", "rating");

dhoro modifiedUrl = existingUrl.Origin + existingUrl.Pathname + 
  "?" + url_query_toString(existingParams);
dekho("Modified:", modifiedUrl);`}
                </code>
              </pre>
            </div>
          </div>
        </div>
      </section>

      {/* Best Practices */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6">Best Practices</h2>
        
        <div className="grid md:grid-cols-2 gap-6">
          <div className="border-l-4 border-green-500 bg-green-50 dark:bg-green-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-green-800 dark:text-green-300">
              ✅ DO: Always Validate URLs
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Check that URLs are valid before parsing. Handle parsing errors gracefully and provide meaningful error messages to users.
            </p>
          </div>

          <div className="border-l-4 border-green-500 bg-green-50 dark:bg-green-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-green-800 dark:text-green-300">
              ✅ DO: Use url_query_has() Before Getting
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Always check if a parameter exists before getting its value. This prevents null reference errors and makes your code more robust.
            </p>
          </div>

          <div className="border-l-4 border-green-500 bg-green-50 dark:bg-green-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-green-800 dark:text-green-300">
              ✅ DO: Let the API Handle Encoding
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              The API automatically encodes special characters in query parameters. Don't manually encode - let url_query_set() and url_query_append() handle it.
            </p>
          </div>

          <div className="border-l-4 border-green-500 bg-green-50 dark:bg-green-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-green-800 dark:text-green-300">
              ✅ DO: Use Append for Multiple Values
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Use url_query_append() when you need multiple values for the same parameter (e.g., tags, filters). Use url_query_set() to replace all values.
            </p>
          </div>

          <div className="border-l-4 border-red-500 bg-red-50 dark:bg-red-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-red-800 dark:text-red-300">
              ❌ DON'T: Manually Build Query Strings
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Never concatenate query strings manually with "&" and "=". Use URLSearchParams to ensure proper encoding and avoid injection vulnerabilities.
            </p>
          </div>

          <div className="border-l-4 border-red-500 bg-red-50 dark:bg-red-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-red-800 dark:text-red-300">
              ❌ DON'T: Assume Parameter Existence
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Don't assume query parameters exist without checking. url_query_get() returns null for missing parameters - always handle this case.
            </p>
          </div>

          <div className="border-l-4 border-red-500 bg-red-50 dark:bg-red-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-red-800 dark:text-red-300">
              ❌ DON'T: Trust User Input URLs
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Always validate and sanitize URLs from user input. Check the protocol, hostname, and parameters to prevent security issues like SSRF.
            </p>
          </div>

          <div className="border-l-4 border-red-500 bg-red-50 dark:bg-red-900/20 p-6 rounded">
            <h3 className="text-xl font-semibold mb-3 text-red-800 dark:text-red-300">
              ❌ DON'T: Modify URL Strings Directly
            </h3>
            <p className="text-gray-700 dark:text-gray-300 text-sm">
              Don't use string manipulation to modify URLs. Parse the URL, modify the URLSearchParams, then reconstruct - this ensures correctness.
            </p>
          </div>
        </div>
      </section>

      {/* Performance Tips */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6">Performance Tips</h2>
        
        <div className="bg-gradient-to-r from-yellow-50 to-orange-50 dark:from-yellow-900/20 dark:to-orange-900/20 rounded-lg p-6">
          <div className="space-y-4">
            <div className="flex items-start space-x-3">
              <span className="text-2xl">🚀</span>
              <div>
                <h4 className="font-semibold mb-1">Cache Parsed URLs</h4>
                <p className="text-sm text-gray-700 dark:text-gray-300">
                  If you're parsing the same URL multiple times, cache the parsed result. URL parsing is relatively fast but caching can improve performance in tight loops.
                </p>
              </div>
            </div>

            <div className="flex items-start space-x-3">
              <span className="text-2xl">💾</span>
              <div>
                <h4 className="font-semibold mb-1">Reuse URLSearchParams Objects</h4>
                <p className="text-sm text-gray-700 dark:text-gray-300">
                  Create a URLSearchParams object once and modify it as needed rather than creating new ones. This is more efficient for building multiple similar URLs.
                </p>
              </div>
            </div>

            <div className="flex items-start space-x-3">
              <span className="text-2xl">⚡</span>
              <div>
                <h4 className="font-semibold mb-1">Batch Parameter Operations</h4>
                <p className="text-sm text-gray-700 dark:text-gray-300">
                  When setting multiple parameters, do all operations before calling url_query_toString(). Convert to string only once at the end.
                </p>
              </div>
            </div>

            <div className="flex items-start space-x-3">
              <span className="text-2xl">🔄</span>
              <div>
                <h4 className="font-semibold mb-1">Avoid Redundant Parsing</h4>
                <p className="text-sm text-gray-700 dark:text-gray-300">
                  Don't parse a URL, convert it to string, and parse it again. Keep the parsed objects in memory and only convert to strings when needed for HTTP requests.
                </p>
              </div>
            </div>

            <div className="flex items-start space-x-3">
              <span className="text-2xl">📦</span>
              <div>
                <h4 className="font-semibold mb-1">Minimize Query String Size</h4>
                <p className="text-sm text-gray-700 dark:text-gray-300">
                  Keep query strings concise. Use short parameter names and remove unnecessary parameters. Large query strings can impact performance and may hit browser/server limits.
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Common Patterns */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6">Common Patterns</h2>

        <div className="space-y-6">
          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-xl font-semibold mb-3">Pattern 1: URL Builder Helper</h3>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Reusable URL builder
kaj URLBuilder(baseUrl) {
  dhoro params = url_query_params("");
  
  ferao {
    setParam: kaj(key, value) {
      url_query_set(params, key, value);
      ferao ei;  // Return self for chaining
    },
    
    appendParam: kaj(key, value) {
      url_query_append(params, key, value);
      ferao ei;
    },
    
    removeParam: kaj(key) {
      url_query_delete(params, key);
      ferao ei;
    },
    
    build: kaj() {
      dhoro queryString = url_query_toString(params);
      ferao queryString != "" ? baseUrl + "?" + queryString : baseUrl;
    }
  };
}

// Usage with method chaining
dhoro url = URLBuilder("https://api.example.com/users")
  .setParam("status", "active")
  .setParam("role", "admin")
  .setParam("page", "1")
  .appendParam("tag", "verified")
  .appendParam("tag", "premium")
  .build();

dekho(url);
// "https://api.example.com/users?status=active&role=admin&page=1&tag=verified&tag=premium"`}
                </code>
              </pre>
            </div>
          </div>

          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-xl font-semibold mb-3">Pattern 2: Query Parameter Merger</h3>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Merge query parameters from multiple sources
kaj mergeQueryParams(params1, params2) {
  dhoro merged = url_query_params("");
  
  // Add all from first params
  dhoro keys1 = url_query_keys(params1);
  dhoro values1 = url_query_values(params1);
  ghuriye (dhoro i = 0; i < dorghyo(keys1); i = i + 1) {
    url_query_append(merged, keys1[i], values1[i]);
  }
  
  // Add all from second params (may override)
  dhoro keys2 = url_query_keys(params2);
  dhoro values2 = url_query_values(params2);
  ghuriye (dhoro i = 0; i < dorghyo(keys2); i = i + 1) {
    // Use set to override, or append to add multiple
    url_query_set(merged, keys2[i], values2[i]);
  }
  
  ferao merged;
}

// Usage
dhoro defaultParams = url_query_params("?limit=20&sort=asc");
dhoro userParams = url_query_params("?page=5&filter=active");
dhoro finalParams = mergeQueryParams(defaultParams, userParams);

dekho(url_query_toString(finalParams));
// "limit=20&sort=asc&page=5&filter=active"`}
                </code>
              </pre>
            </div>
          </div>

          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-xl font-semibold mb-3">Pattern 3: URL Router</h3>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Route handler based on URL path and params
kaj routeRequest(urlString) {
  dhoro url = url_parse(urlString);
  dhoro path = url.Pathname;
  dhoro params = url_query_params(url.Search);
  
  // Route based on path
  jodi (path == "/api/users") {
    // Handle users endpoint
    dhoro page = url_query_has(params, "page") ? 
      url_query_get(params, "page") : "1";
    dhoro limit = url_query_has(params, "limit") ? 
      url_query_get(params, "limit") : "20";
    
    ferao handleUsers(page, limit, params);
    
  } nahole jodi (path == "/api/products") {
    // Handle products endpoint
    dhoro category = url_query_get(params, "category");
    dhoro minPrice = url_query_get(params, "minPrice");
    dhoro maxPrice = url_query_get(params, "maxPrice");
    
    ferao handleProducts(category, minPrice, maxPrice);
    
  } nahole jodi (khuje(path, "/api/users/") == 0) {
    // Handle specific user by ID
    dhoro pathParts = bibhajan(path, "/");
    dhoro userId = pathParts[3];
    
    ferao handleUserById(userId, params);
  }
  
  ferao {status: 404, message: "Route not found"};
}

// Helper functions
kaj handleUsers(page, limit, params) {
  dhoro filters = [];
  
  jodi (url_query_has(params, "status")) {
    joro(filters, "status=" + url_query_get(params, "status"));
  }
  
  jodi (url_query_has(params, "role")) {
    joro(filters, "role=" + url_query_get(params, "role"));
  }
  
  dekho("Fetching users: page=" + page + ", limit=" + limit);
  dekho("Filters:", joro(filters, ", "));
  
  ferao {status: 200, data: "users list"};
}

// Usage
dhoro result = routeRequest("https://api.com/api/users?page=2&status=active&role=admin");
dekho(result);`}
                </code>
              </pre>
            </div>
          </div>

          <div className="border rounded-lg p-6 bg-white dark:bg-gray-800">
            <h3 className="text-xl font-semibold mb-3">Pattern 4: Pagination Helper</h3>
            <div className="bg-gray-50 dark:bg-gray-900 rounded p-4">
              <pre className="text-sm overflow-x-auto">
                <code className="language-banglacode">
{`// Build paginated URLs
kaj buildPaginationURLs(currentUrl, totalPages) {
  dhoro url = url_parse(currentUrl);
  dhoro params = url_query_params(url.Search);
  
  dhoro currentPage = url_query_has(params, "page") ? 
    text_shongkha(url_query_get(params, "page")) : 1;
  
  dhoro baseUrl = url.Origin + url.Pathname;
  
  // Build URL for specific page
  dhoro buildPageUrl = kaj(pageNum) {
    dhoro newParams = url_query_params(url.Search);
    url_query_set(newParams, "page", shongkha_text(pageNum));
    ferao baseUrl + "?" + url_query_toString(newParams);
  };
  
  dhoro urls = {
    first: buildPageUrl(1),
    prev: currentPage > 1 ? buildPageUrl(currentPage - 1) : mittha,
    current: currentUrl,
    next: currentPage < totalPages ? buildPageUrl(currentPage + 1) : mittha,
    last: buildPageUrl(totalPages)
  };
  
  ferao urls;
}

// Usage
dhoro currentUrl = "https://api.com/users?status=active&page=3&limit=20";
dhoro pagination = buildPaginationURLs(currentUrl, 10);

dekho("First:", pagination.first);
dekho("Previous:", pagination.prev);
dekho("Next:", pagination.next);
dekho("Last:", pagination.last);`}
                </code>
              </pre>
            </div>
          </div>
        </div>
      </section>

      {/* Related APIs */}
      <section className="mb-12">
        <h2 className="text-3xl font-bold mb-6">Related APIs</h2>
        
        <div className="grid md:grid-cols-3 gap-6">
          <a href="/docs/http" className="border rounded-lg p-6 bg-white dark:bg-gray-800 hover:shadow-lg transition-shadow">
            <h3 className="text-xl font-semibold mb-2 text-blue-600 dark:text-blue-400">HTTP API</h3>
            <p className="text-sm text-gray-700 dark:text-gray-300">
              Make HTTP requests with the URLs you build. Perfect companion for API interactions.
            </p>
          </a>

          <a href="/docs/networking" className="border rounded-lg p-6 bg-white dark:bg-gray-800 hover:shadow-lg transition-shadow">
            <h3 className="text-xl font-semibold mb-2 text-green-600 dark:text-green-400">Networking</h3>
            <p className="text-sm text-gray-700 dark:text-gray-300">
              Low-level networking APIs for TCP/UDP connections and custom protocols.
            </p>
          </a>

          <a href="/docs/json" className="border rounded-lg p-6 bg-white dark:bg-gray-800 hover:shadow-lg transition-shadow">
            <h3 className="text-xl font-semibold mb-2 text-purple-600 dark:text-purple-400">JSON API</h3>
            <p className="text-sm text-gray-700 dark:text-gray-300">
              Parse and serialize JSON data received from API requests.
            </p>
          </a>
        </div>
      </section>

      {/* Summary */}
      <section className="bg-gradient-to-r from-blue-50 to-indigo-50 dark:from-blue-900/20 dark:to-indigo-900/20 rounded-lg p-8">
        <h2 className="text-3xl font-bold mb-4">Summary</h2>
        <div className="space-y-4 text-gray-700 dark:text-gray-300">
          <p>
            The URL Parsing API in BanglaCode provides comprehensive tools for working with URLs and query parameters. Whether you're building web APIs, parsing external URLs, or constructing complex query strings, these functions make URL operations safe, reliable, and easy.
          </p>
          <p>
            Key benefits:
          </p>
          <ul className="list-disc list-inside space-y-2 ml-4">
            <li><strong>Complete URL Parsing:</strong> Extract all components from any URL string</li>
            <li><strong>Safe Query Handling:</strong> Automatic encoding/decoding prevents injection attacks</li>
            <li><strong>Flexible Parameters:</strong> Support for single and multiple values per key</li>
            <li><strong>Easy Manipulation:</strong> Simple API for getting, setting, and deleting parameters</li>
            <li><strong>API Integration:</strong> Perfect for building REST API clients and web services</li>
          </ul>
          <p>
            Use the URL Parsing API whenever you need to work with web addresses, build API requests, handle query parameters, or route requests based on URL patterns. Combined with the HTTP API, it provides everything you need for web and API development.
          </p>
        </div>
      </section>
    </div>
  );
}
