export default function BufferPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold mb-2">Buffer API (বাফার)</h1>
        <p className="text-lg text-muted-foreground">
          Binary data handling for efficient manipulation of raw bytes, file I/O, network protocols, and encoding conversions.
        </p>
      </div>

      {/* Overview */}
      <section className="space-y-3">
        <h2 className="text-2xl font-semibold border-b pb-2">Overview</h2>
        <p>
          The Buffer API provides powerful tools for working with binary data in BanglaCode. Unlike strings that handle text,
          Buffers work with raw bytes, making them essential for:
        </p>
        <ul className="list-disc list-inside space-y-1 ml-4">
          <li><strong>File I/O:</strong> Reading and writing binary files (images, audio, video)</li>
          <li><strong>Network protocols:</strong> Building custom binary protocols (headers, payloads, checksums)</li>
          <li><strong>Encoding conversions:</strong> Converting between UTF-8, hex, base64</li>
          <li><strong>Performance:</strong> Efficient memory usage for large data processing</li>
        </ul>
        <div className="bg-blue-50 dark:bg-blue-950 p-4 rounded-lg border border-blue-200 dark:border-blue-800 mt-4">
          <p className="text-sm">
            <strong>💡 Key Concept:</strong> Buffers are thread-safe and immutable by design. Operations that modify buffers
            return new Buffer instances, preventing accidental data corruption in concurrent code.
          </p>
        </div>
      </section>

      {/* Quick Start */}
      <section className="space-y-3">
        <h2 className="text-2xl font-semibold border-b pb-2">Quick Start</h2>
        <div className="bg-gray-50 dark:bg-gray-900 p-4 rounded-lg">
          <pre className="text-sm overflow-x-auto">
            <code>{`// Create a buffer with size
dhoro buf = buffer_banao(10);  // 10-byte buffer

// Create buffer from string
dhoro textBuf = buffer_theke("Hello World");

// Create buffer from byte array
dhoro byteBuf = buffer_theke([72, 101, 108, 108, 111]);

// Convert buffer to string
dhoro text = buffer_text(byteBuf);
dekho(text);  // "Hello"

// Concatenate buffers
dhoro combined = buffer_joro(textBuf, byteBuf);

// Slice buffer
dhoro slice = buffer_angsho(combined, 0, 5);`}</code>
          </pre>
        </div>
      </section>

      {/* API Reference */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold border-b pb-2">API Reference</h2>

        {/* buffer_banao */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            buffer_banao(size)
          </h3>
          <p className="text-sm text-muted-foreground">বাফার বানাও - Create buffer</p>
          <p>Creates a new Buffer with the specified size in bytes, initialized with zeros.</p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`dhoro buf = buffer_banao(1024);  // 1KB buffer filled with zeros`}</code>
            </pre>
          </div>
          <div className="text-sm">
            <strong>Parameters:</strong>
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>size</code> (Number) - Size in bytes (must be positive)</li>
            </ul>
            <strong className="block mt-2">Returns:</strong> Buffer object
          </div>
        </div>

        {/* buffer_theke */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            buffer_theke(data)
          </h3>
          <p className="text-sm text-muted-foreground">বাফার থেকে - Buffer from</p>
          <p>Creates a Buffer from a string or byte array. Supports UTF-8 encoding for strings.</p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`// From string
dhoro buf1 = buffer_theke("Hello");

// From byte array
dhoro buf2 = buffer_theke([72, 101, 108, 108, 111]);

// Both create identical buffers representing "Hello"`}</code>
            </pre>
          </div>
          <div className="text-sm">
            <strong>Parameters:</strong>
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>data</code> (String | Array) - Source data (string or byte array)</li>
            </ul>
            <strong className="block mt-2">Returns:</strong> Buffer object
          </div>
        </div>

        {/* buffer_joro */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            buffer_joro(buf1, buf2, ...)
          </h3>
          <p className="text-sm text-muted-foreground">বাফার জোড়ো - Join buffers</p>
          <p>Concatenates multiple buffers into a single buffer. Accepts variable number of arguments.</p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`dhoro header = buffer_theke([255, 1]);
dhoro body = buffer_theke("Data");
dhoro footer = buffer_theke([171, 205]);

dhoro message = buffer_joro(header, body, footer);
// Result: [255, 1, 68, 97, 116, 97, 171, 205]`}</code>
            </pre>
          </div>
          <div className="text-sm">
            <strong>Parameters:</strong>
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>...buffers</code> (Buffer) - Buffers to concatenate (at least 1)</li>
            </ul>
            <strong className="block mt-2">Returns:</strong> New combined Buffer
          </div>
        </div>

        {/* buffer_text */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            buffer_text(buf, encoding?)
          </h3>
          <p className="text-sm text-muted-foreground">বাফার টেক্সট - Buffer to text</p>
          <p>Converts a Buffer to a string using the specified encoding. Defaults to UTF-8.</p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`dhoro buf = buffer_theke([72, 101, 108, 108, 111]);

dhoro utf8 = buffer_text(buf);           // "Hello"
dhoro utf8Explicit = buffer_text(buf, "utf8");  // "Hello"
dhoro hex = buffer_text(buf, "hex");     // "48656c6c6f"`}</code>
            </pre>
          </div>
          <div className="text-sm">
            <strong>Parameters:</strong>
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>buf</code> (Buffer) - Buffer to convert</li>
              <li><code>encoding</code> (String, optional) - Encoding type ("utf8", "hex", "base64")</li>
            </ul>
            <strong className="block mt-2">Returns:</strong> String representation
          </div>
        </div>

        {/* buffer_lekho */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            buffer_lekho(buf, data, offset)
          </h3>
          <p className="text-sm text-muted-foreground">বাফার লেখো - Write to buffer</p>
          <p>Writes string data to a buffer at the specified offset. Returns number of bytes written.</p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`dhoro buf = buffer_banao(20);

buffer_lekho(buf, "Hello", 0);   // Write at start
buffer_lekho(buf, "World", 10);  // Write at offset 10

dhoro text = buffer_text(buf);   // "Hello\\x00\\x00\\x00\\x00\\x00World..."`}</code>
            </pre>
          </div>
          <div className="text-sm">
            <strong>Parameters:</strong>
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>buf</code> (Buffer) - Target buffer (modified in place)</li>
              <li><code>data</code> (String) - String to write</li>
              <li><code>offset</code> (Number) - Starting position (must be within buffer bounds)</li>
            </ul>
            <strong className="block mt-2">Returns:</strong> Number of bytes written
          </div>
        </div>

        {/* buffer_angsho */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            buffer_angsho(buf, start, end?)
          </h3>
          <p className="text-sm text-muted-foreground">বাফার অংশ - Buffer slice</p>
          <p>
            Extracts a section of the buffer. If <code>end</code> is omitted, slices from <code>start</code> to end of buffer.
          </p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`dhoro buf = buffer_theke("Hello World");

dhoro hello = buffer_angsho(buf, 0, 5);   // "Hello"
dhoro world = buffer_angsho(buf, 6, 11);  // "World"
dhoro tail = buffer_angsho(buf, 6);       // "World" (from 6 to end)`}</code>
            </pre>
          </div>
          <div className="text-sm">
            <strong>Parameters:</strong>
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>buf</code> (Buffer) - Source buffer</li>
              <li><code>start</code> (Number) - Starting index (inclusive)</li>
              <li><code>end</code> (Number, optional) - Ending index (exclusive, defaults to buffer length)</li>
            </ul>
            <strong className="block mt-2">Returns:</strong> New Buffer containing the slice
          </div>
        </div>

        {/* buffer_tulona */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            buffer_tulona(buf1, buf2)
          </h3>
          <p className="text-sm text-muted-foreground">বাফার তুলনা - Compare buffers</p>
          <p>Compares two buffers lexicographically. Returns -1, 0, or 1.</p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`dhoro buf1 = buffer_theke("ABC");
dhoro buf2 = buffer_theke("ABC");
dhoro buf3 = buffer_theke("DEF");

buffer_tulona(buf1, buf2);  // 0 (equal)
buffer_tulona(buf1, buf3);  // -1 (buf1 < buf3)
buffer_tulona(buf3, buf1);  // 1 (buf3 > buf1)`}</code>
            </pre>
          </div>
          <div className="text-sm">
            <strong>Parameters:</strong>
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>buf1</code> (Buffer) - First buffer</li>
              <li><code>buf2</code> (Buffer) - Second buffer</li>
            </ul>
            <strong className="block mt-2">Returns:</strong> Number (-1, 0, or 1)
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>-1</code> if buf1 &lt; buf2</li>
              <li><code>0</code> if buf1 == buf2</li>
              <li><code>1</code> if buf1 &gt; buf2</li>
            </ul>
          </div>
        </div>

        {/* buffer_hex */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            buffer_hex(buf)
          </h3>
          <p className="text-sm text-muted-foreground">বাফার হেক্স - Buffer to hex</p>
          <p>Converts a buffer to hexadecimal string representation (lowercase).</p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`dhoro buf = buffer_theke([255, 0, 127, 16]);
dhoro hex = buffer_hex(buf);
dekho(hex);  // "ff007f10"`}</code>
            </pre>
          </div>
          <div className="text-sm">
            <strong>Parameters:</strong>
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>buf</code> (Buffer) - Buffer to convert</li>
            </ul>
            <strong className="block mt-2">Returns:</strong> Hexadecimal string
          </div>
        </div>

        {/* buffer_copy */}
        <div className="border rounded-lg p-4 space-y-2">
          <h3 className="text-xl font-semibold text-blue-600 dark:text-blue-400">
            buffer_copy(target, source, targetStart)
          </h3>
          <p className="text-sm text-muted-foreground">বাফার কপি - Copy buffer</p>
          <p>Copies data from source buffer to target buffer starting at the specified offset.</p>
          <div className="bg-gray-50 dark:bg-gray-900 p-3 rounded">
            <pre className="text-sm">
              <code>{`dhoro target = buffer_banao(20);
dhoro source = buffer_theke("Hello");

dhoro written = buffer_copy(target, source, 0);
dekho(written);  // 5 (bytes copied)

dhoro text = buffer_text(target);
dekho(text);  // "Hello\\x00\\x00\\x00\\x00..."`}</code>
            </pre>
          </div>
          <div className="text-sm">
            <strong>Parameters:</strong>
            <ul className="list-disc list-inside ml-4 mt-1">
              <li><code>target</code> (Buffer) - Destination buffer (modified in place)</li>
              <li><code>source</code> (Buffer) - Source buffer</li>
              <li><code>targetStart</code> (Number) - Starting offset in target</li>
            </ul>
            <strong className="block mt-2">Returns:</strong> Number of bytes copied
          </div>
        </div>
      </section>

      {/* Real-World Examples */}
      <section className="space-y-4">
        <h2 className="text-2xl font-semibold border-b pb-2">Real-World Examples</h2>

        {/* Example 1: Binary Protocol */}
        <div className="space-y-2">
          <h3 className="text-lg font-semibold">Example 1: Binary Network Protocol</h3>
          <p className="text-sm text-muted-foreground">
            Build a custom binary protocol message with header, payload, and checksum.
          </p>
          <div className="bg-gray-50 dark:bg-gray-900 p-4 rounded-lg">
            <pre className="text-sm overflow-x-auto">
              <code>{`// Define protocol structure:
// - Header: 3 bytes (magic 255, version, message type)
// - Payload: variable length (UTF-8 string)
// - Checksum: 2 bytes

// Build message
dhoro header = buffer_theke([255, 1, 2]);  // Magic=255, Version=1, Type=2
dhoro payload = buffer_theke("Important data");
dhoro checksum = buffer_theke([171, 205]);  // Simple checksum

// Combine all parts
dhoro message = buffer_joro(header, payload, checksum);
dekho("Total message size:", dorghyo(buffer_text(message)), "bytes");

// Parse received message
dhoro receivedHeader = buffer_angsho(message, 0, 3);
dhoro receivedPayload = buffer_angsho(message, 3, 17);
dhoro receivedChecksum = buffer_angsho(message, 17);

// Verify magic number
dhoro magicByte = buffer_angsho(receivedHeader, 0, 1);
jodi (buffer_tulona(magicByte, buffer_theke([255])) == 0) {
    dekho("Valid protocol message");
    dekho("Payload:", buffer_text(receivedPayload));
} nahole {
    dekho("Invalid magic number");
}`}</code>
            </pre>
          </div>
        </div>

        {/* Example 2: File Encoding Conversion */}
        <div className="space-y-2">
          <h3 className="text-lg font-semibold">Example 2: File Encoding Conversion</h3>
          <p className="text-sm text-muted-foreground">
            Read a file, convert its encoding, and write it back.
          </p>
          <div className="bg-gray-50 dark:bg-gray-900 p-4 rounded-lg">
            <pre className="text-sm overflow-x-auto">
              <code>{`// Read file as binary
dhoro fileContent = poro("input.txt");
dhoro buf = buffer_theke(fileContent);

// Convert to hex for inspection or transmission
dhoro hexContent = buffer_hex(buf);
dekho("Hex representation:", hexContent);

// Write hex to file for debugging
lekho("output.hex", hexContent);

// Convert back to UTF-8 and process
dhoro text = buffer_text(buf, "utf8");
dekho("Original text:", text);`}</code>
            </pre>
          </div>
        </div>

        {/* Example 3: Memory-Efficient Data Processing */}
        <div className="space-y-2">
          <h3 className="text-lg font-semibold">Example 3: Efficient Binary Data Processing</h3>
          <p className="text-sm text-muted-foreground">
            Process large binary data in chunks to optimize memory usage.
          </p>
          <div className="bg-gray-50 dark:bg-gray-900 p-4 rounded-lg">
            <pre className="text-sm overflow-x-auto">
              <code>{`// Simulate reading a large file in chunks
dhoro largeData = buffer_theke("This is a very long data stream that needs processing...");
dhoro chunkSize = 10;
dhoro totalChunks = 0;

// Process in chunks
dhoro offset = 0;
jotokkhon (offset < dorghyo(buffer_text(largeData))) {
    // Extract chunk
    dhoro chunk = buffer_angsho(largeData, offset, offset + chunkSize);
    
    // Process chunk (e.g., search for pattern, transform data)
    dhoro chunkText = buffer_text(chunk);
    dekho("Processing chunk", totalChunks + 1, ":", chunkText);
    
    offset = offset + chunkSize;
    totalChunks = totalChunks + 1;
}

dekho("Total chunks processed:", totalChunks);`}</code>
            </pre>
          </div>
        </div>

        {/* Example 4: Byte Manipulation */}
        <div className="space-y-2">
          <h3 className="text-lg font-semibold">Example 4: Low-Level Byte Manipulation</h3>
          <p className="text-sm text-muted-foreground">
            Direct byte-level operations for performance-critical code.
          </p>
          <div className="bg-gray-50 dark:bg-gray-900 p-4 rounded-lg">
            <pre className="text-sm overflow-x-auto">
              <code>{`// Create buffer with specific bytes
dhoro buf = buffer_banao(8);

// Write integers as bytes
buffer_lekho(buf, "AB", 0);  // First 2 bytes
buffer_lekho(buf, "CD", 4);  // Bytes 4-5

// Read specific sections
dhoro part1 = buffer_angsho(buf, 0, 2);
dhoro part2 = buffer_angsho(buf, 4, 6);

dekho("Part 1:", buffer_text(part1));
dekho("Part 2:", buffer_text(part2));

// Compare parts
jodi (buffer_tulona(part1, part2) != 0) {
    dekho("Parts are different");
}`}</code>
            </pre>
          </div>
        </div>
      </section>

      {/* Best Practices */}
      <section className="space-y-3">
        <h2 className="text-2xl font-semibold border-b pb-2">Best Practices</h2>
        
        <div className="space-y-4">
          <div className="bg-green-50 dark:bg-green-950 p-4 rounded-lg border border-green-200 dark:border-green-800">
            <h3 className="font-semibold text-green-800 dark:text-green-200 mb-2">✅ DO:</h3>
            <ul className="list-disc list-inside space-y-1 text-sm">
              <li><strong>Pre-allocate buffers:</strong> Use <code>buffer_banao()</code> with known size for better performance</li>
              <li><strong>Reuse buffers:</strong> When processing streams, reuse buffer instances to reduce allocations</li>
              <li><strong>Use slicing for parsing:</strong> <code>buffer_angsho()</code> is efficient for extracting protocol parts</li>
              <li><strong>Validate buffer sizes:</strong> Always check bounds before writing to prevent overflows</li>
              <li><strong>Choose appropriate encoding:</strong> Use hex for debugging, UTF-8 for text, binary for performance</li>
            </ul>
          </div>

          <div className="bg-red-50 dark:bg-red-950 p-4 rounded-lg border border-red-200 dark:border-red-800">
            <h3 className="font-semibold text-red-800 dark:text-red-200 mb-2">❌ DON'T:</h3>
            <ul className="list-disc list-inside space-y-1 text-sm">
              <li><strong>Don't concatenate repeatedly:</strong> Use <code>buffer_joro()</code> with all buffers at once, not in a loop</li>
              <li><strong>Don't use buffers for small strings:</strong> Regular strings are more efficient for text-only data</li>
              <li><strong>Don't ignore encoding:</strong> Always specify encoding when converting to/from strings</li>
              <li><strong>Don't assume immutability:</strong> <code>buffer_lekho()</code> and <code>buffer_copy()</code> modify buffers in place</li>
              <li><strong>Don't forget to validate:</strong> Check buffer sizes before operations to prevent runtime errors</li>
            </ul>
          </div>
        </div>
      </section>

      {/* Performance Tips */}
      <section className="space-y-3">
        <h2 className="text-2xl font-semibold border-b pb-2">Performance Considerations</h2>
        
        <div className="bg-yellow-50 dark:bg-yellow-950 p-4 rounded-lg border border-yellow-200 dark:border-yellow-800">
          <h3 className="font-semibold mb-2">⚡ Optimization Tips:</h3>
          <ul className="list-disc list-inside space-y-1 text-sm">
            <li>
              <strong>Batch operations:</strong> Combine multiple small buffers into one with <code>buffer_joro()</code> instead of 
              multiple concatenations
            </li>
            <li>
              <strong>Buffer pooling:</strong> For high-throughput applications, reuse buffer instances instead of creating new ones
            </li>
            <li>
              <strong>Avoid conversions:</strong> Keep data in buffer form as long as possible; convert to string only when needed
            </li>
            <li>
              <strong>Use appropriate sizes:</strong> Allocate buffers with the exact size needed to avoid wasted memory
            </li>
            <li>
              <strong>Slicing is cheap:</strong> <code>buffer_angsho()</code> creates a view, making it efficient for parsing large data
            </li>
          </ul>
        </div>
      </section>

      {/* Common Use Cases */}
      <section className="space-y-3">
        <h2 className="text-2xl font-semibold border-b pb-2">Common Use Cases</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div className="border rounded-lg p-4 space-y-2">
            <h3 className="font-semibold text-blue-600 dark:text-blue-400">Binary File I/O</h3>
            <p className="text-sm">
              Read/write images, audio, video, or any binary format efficiently without text conversion overhead.
            </p>
          </div>
          <div className="border rounded-lg p-4 space-y-2">
            <h3 className="font-semibold text-blue-600 dark:text-blue-400">Network Protocols</h3>
            <p className="text-sm">
              Build custom binary protocols with headers, payloads, and checksums for efficient network communication.
            </p>
          </div>
          <div className="border rounded-lg p-4 space-y-2">
            <h3 className="font-semibold text-blue-600 dark:text-blue-400">Encoding Conversions</h3>
            <p className="text-sm">
              Convert between UTF-8, hex, and base64 encodings for data transmission, storage, or debugging.
            </p>
          </div>
          <div className="border rounded-lg p-4 space-y-2">
            <h3 className="font-semibold text-blue-600 dark:text-blue-400">Cryptography</h3>
            <p className="text-sm">
              Work with raw bytes for encryption, hashing, and digital signatures where binary precision is critical.
            </p>
          </div>
        </div>
      </section>

      {/* Related Features */}
      <section className="space-y-3">
        <h2 className="text-2xl font-semibold border-b pb-2">Related Features</h2>
        <div className="flex flex-wrap gap-2">
          <a href="/docs/functions" className="px-3 py-1 bg-blue-100 dark:bg-blue-900 text-blue-800 dark:text-blue-200 rounded-full text-sm hover:bg-blue-200 dark:hover:bg-blue-800 transition-colors">
            File I/O Functions
          </a>
          <a href="/docs/networking" className="px-3 py-1 bg-blue-100 dark:bg-blue-900 text-blue-800 dark:text-blue-200 rounded-full text-sm hover:bg-blue-200 dark:hover:bg-blue-800 transition-colors">
            Networking
          </a>
          <a href="/docs/async" className="px-3 py-1 bg-blue-100 dark:bg-blue-900 text-blue-800 dark:text-blue-200 rounded-full text-sm hover:bg-blue-200 dark:hover:bg-blue-800 transition-colors">
            Async/Await
          </a>
        </div>
      </section>

      {/* Summary */}
      <section className="bg-blue-50 dark:bg-blue-950 p-6 rounded-lg border border-blue-200 dark:border-blue-800">
        <h2 className="text-xl font-semibold mb-3">Summary</h2>
        <p className="text-sm mb-3">
          The Buffer API provides powerful, performance-optimized tools for working with binary data in BanglaCode. 
          Whether you're building network protocols, processing binary files, or handling encoding conversions, 
          Buffers offer the low-level control and efficiency you need.
        </p>
        <p className="text-sm">
          <strong>Key takeaways:</strong> Use <code>buffer_banao()</code> for pre-allocation, <code>buffer_joro()</code> 
          for combining data, and <code>buffer_angsho()</code> for efficient parsing. Always validate buffer sizes 
          and choose appropriate encodings for your use case.
        </p>
      </section>
    </div>
  );
}
