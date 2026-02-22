import CodeBlock from "@/components/CodeBlock";
import DocNavigation from "@/components/DocNavigation";

export default function EventEmitterPage() {
  return (
    <div>
      <div className="flex items-center gap-2 text-sm text-muted-foreground mb-4">
        <span className="px-2 py-1 bg-primary/10 text-primary rounded-full text-xs font-medium">
          Advanced
        </span>
      </div>

      <h1>EventEmitter - Event-Driven Architecture</h1>

      <p className="lead text-xl text-muted-foreground mt-4">
        BanglaCode supports event-driven programming through the <code>EventEmitter</code> API. This allows you to build reactive applications where components communicate through events, perfect for building message buses, pub/sub systems, and event-driven architectures.
      </p>

      <div className="bg-blue-50 dark:bg-blue-950 border border-blue-200 dark:border-blue-800 rounded-lg p-4 my-6">
        <h4 className="text-blue-900 dark:text-blue-100 font-semibold mb-2">Key Concepts</h4>
        <ul className="text-blue-800 dark:text-blue-200 space-y-1">
          <li><code>ghotona_srishti()</code> (ঘটনা সৃষ্টি) - Create an EventEmitter instance</li>
          <li><code>ghotona_shuno()</code> (ঘটনা শুনো) - Listen to events</li>
          <li><code>ghotona_ekbar()</code> (ঘটনা একবার) - Listen to event once</li>
          <li><code>ghotona_prokash()</code> (ঘটনা প্রকাশ) - Emit/trigger events</li>
          <li><code>ghotona_bondho()</code> (ঘটনা বন্ধ) - Remove listener</li>
        </ul>
      </div>

      <h2>Creating an EventEmitter</h2>

      <p>
        Create a new EventEmitter instance using <code>ghotona_srishti()</code>:
      </p>

      <CodeBlock
        code={`// Create an event emitter
dhoro emitter = ghotona_srishti();
dekho(emitter);  // EventEmitter(events=0, listeners=0)`}
      />

      <h2>Adding Event Listeners</h2>

      <h3>ghotona_shuno() - On</h3>

      <p>
        Use <code>ghotona_shuno()</code> to add an event listener. The listener will be called every time the event is emitted:
      </p>

      <CodeBlock
        code={`dhoro emitter = ghotona_srishti();

// Add a listener
ghotona_shuno(emitter, "message", kaj(data) {
    dekho("Received:", data);
});

// Emit the event multiple times
ghotona_prokash(emitter, "message", "Hello");  // Received: Hello
ghotona_prokash(emitter, "message", "World");  // Received: World`}
      />

      <h3>ghotona_ekbar() - Once</h3>

      <p>
        Use <code>ghotona_ekbar()</code> to add a listener that only runs once:
      </p>

      <CodeBlock
        code={`dhoro emitter = ghotona_srishti();

// Add a one-time listener
ghotona_ekbar(emitter, "startup", kaj() {
    dekho("Application started!");
});

// Emit multiple times - listener only runs once
ghotona_prokash(emitter, "startup");  // Application started!
ghotona_prokash(emitter, "startup");  // (no output - listener already removed)`}
      />

      <h2>Emitting Events</h2>

      <h3>ghotona_prokash() - Emit</h3>

      <p>
        Use <code>ghotona_prokash()</code> to emit an event with optional data:
      </p>

      <CodeBlock
        code={`dhoro emitter = ghotona_srishti();

ghotona_shuno(emitter, "user:login", kaj(username, time) {
    dekho(username + " logged in at " + time);
});

// Emit with multiple data arguments
ghotona_prokash(emitter, "user:login", "Ankan", "09:30 AM");
// Output: Ankan logged in at 09:30 AM`}
      />

      <h2>Multiple Listeners</h2>

      <p>
        You can add multiple listeners to the same event. All listeners will be called in the order they were added:
      </p>

      <CodeBlock
        code={`dhoro emitter = ghotona_srishti();

// First listener
ghotona_shuno(emitter, "data", kaj(value) {
    dekho("Listener 1:", value);
});

// Second listener
ghotona_shuno(emitter, "data", kaj(value) {
    dekho("Listener 2:", value * 2);
});

ghotona_prokash(emitter, "data", 10);
// Output:
// Listener 1: 10
// Listener 2: 20`}
      />

      <h2>Removing Listeners</h2>

      <h3>ghotona_bondho() - Off</h3>

      <p>
        Remove a specific listener using <code>ghotona_bondho()</code>:
      </p>

      <CodeBlock
        code={`dhoro emitter = ghotona_srishti();

// Store listener reference
dhoro listener = kaj(msg) {
    dekho("Received:", msg);
};

ghotona_shuno(emitter, "update", listener);
ghotona_prokash(emitter, "update", "First");  // Received: First

// Remove the listener
ghotona_bondho(emitter, "update", listener);
ghotona_prokash(emitter, "update", "Second");  // (no output)`}
      />

      <h3>ghotona_sob_bondho() - Remove All</h3>

      <p>
        Remove all listeners for a specific event or all events:
      </p>

      <CodeBlock
        code={`dhoro emitter = ghotona_srishti();

ghotona_shuno(emitter, "event1", kaj() { dekho("Event 1"); });
ghotona_shuno(emitter, "event2", kaj() { dekho("Event 2"); });

// Remove all listeners for "event1"
ghotona_sob_bondho(emitter, "event1");

// Remove all listeners for all events
ghotona_sob_bondho(emitter);`}
      />

      <h2>Inspecting EventEmitter</h2>

      <h3>ghotona_shrotara() - Get Listeners</h3>

      <p>
        Get all listeners for a specific event:
      </p>

      <CodeBlock
        code={`dhoro emitter = ghotona_srishti();

ghotona_shuno(emitter, "test", kaj() { });
ghotona_shuno(emitter, "test", kaj() { });

dhoro listeners = ghotona_shrotara(emitter, "test");
dekho("Number of listeners:", dorghyo(listeners));  // 2`}
      />

      <h3>ghotona_naam_sob() - Get Event Names</h3>

      <p>
        Get all event names that have listeners:
      </p>

      <CodeBlock
        code={`dhoro emitter = ghotona_srishti();

ghotona_shuno(emitter, "event1", kaj() { });
ghotona_shuno(emitter, "event2", kaj() { });
ghotona_shuno(emitter, "event3", kaj() { });

dhoro events = ghotona_naam_sob(emitter);
dekho("Events:", events);  // ["event1", "event2", "event3"]`}
      />

      <h2>Method Chaining</h2>

      <p>
        EventEmitter methods return the emitter, allowing method chaining:
      </p>

      <CodeBlock
        code={`dhoro emitter = ghotona_srishti();

ghotona_shuno(emitter, "start", kaj() { dekho("Starting..."); })
ghotona_shuno(emitter, "process", kaj() { dekho("Processing..."); })
ghotona_shuno(emitter, "end", kaj() { dekho("Done!"); });

// Chain emit calls
ghotona_prokash(emitter, "start");
ghotona_prokash(emitter, "process");
ghotona_prokash(emitter, "end");`}
      />

      <h2>Real-World Examples</h2>

      <h3>Example 1: Message Bus</h3>

      <CodeBlock
        code={`// Create a message bus for pub/sub
dhoro messageBus = ghotona_srishti();

// Subscribe to messages
ghotona_shuno(messageBus, "notification", kaj(msg) {
    dekho("[NOTIFICATION]", msg);
});

ghotona_shuno(messageBus, "error", kaj(err) {
    dekho("[ERROR]", err);
});

// Publish messages
ghotona_prokash(messageBus, "notification", "User logged in");
ghotona_prokash(messageBus, "error", "Database connection failed");`}
      />

      <h3>Example 2: Custom Event Handling</h3>

      <CodeBlock
        code={`// Simple task queue with events
dhoro taskQueue = ghotona_srishti();
dhoro tasks = [];

// Listen for task:add event
ghotona_shuno(taskQueue, "task:add", kaj(task) {
    dhokao(tasks, task);
    dekho("Task added:", task["name"]);
    ghotona_prokash(taskQueue, "task:count", dorghyo(tasks));
});

// Listen for task count updates
ghotona_shuno(taskQueue, "task:count", kaj(count) {
    dekho("Total tasks:", count);
});

// Add tasks
ghotona_prokash(taskQueue, "task:add", {"name": "Send email"});
ghotona_prokash(taskQueue, "task:add", {"name": "Process payment"});
ghotona_prokash(taskQueue, "task:add", {"name": "Update database"});`}
      />

      <h3>Example 3: State Management</h3>

      <CodeBlock
        code={`// Simple state manager with events
dhoro stateManager = ghotona_srishti();
dhoro state = {"count": 0};

// Listen for state changes
ghotona_shuno(stateManager, "state:change", kaj(newState) {
    state = newState;
    dekho("State updated:", state);
});

// Listen for specific actions
ghotona_shuno(stateManager, "increment", kaj() {
    state["count"] = state["count"] + 1;
    ghotona_prokash(stateManager, "state:change", state);
});

ghotona_shuno(stateManager, "decrement", kaj() {
    state["count"] = state["count"] - 1;
    ghotona_prokash(stateManager, "state:change", state);
});

// Trigger actions
ghotona_prokash(stateManager, "increment");  // count: 1
ghotona_prokash(stateManager, "increment");  // count: 2
ghotona_prokash(stateManager, "decrement");  // count: 1`}
      />

      <h2>API Reference</h2>

      <div className="space-y-4 my-6">
        <div className="border rounded-lg p-4">
          <h4 className="font-semibold mb-2"><code>ghotona_srishti()</code></h4>
          <p className="text-sm text-muted-foreground mb-2">
            Creates a new EventEmitter instance.
          </p>
          <p className="text-sm"><strong>Returns:</strong> EventEmitter</p>
        </div>

        <div className="border rounded-lg p-4">
          <h4 className="font-semibold mb-2"><code>ghotona_shuno(emitter, event_name, callback)</code></h4>
          <p className="text-sm text-muted-foreground mb-2">
            Adds an event listener that runs every time the event is emitted.
          </p>
          <p className="text-sm"><strong>Parameters:</strong></p>
          <ul className="text-sm list-disc list-inside ml-4">
            <li><code>emitter</code> - EventEmitter instance</li>
            <li><code>event_name</code> - String event name</li>
            <li><code>callback</code> - Function to call when event is emitted</li>
          </ul>
          <p className="text-sm mt-2"><strong>Returns:</strong> EventEmitter (for chaining)</p>
        </div>

        <div className="border rounded-lg p-4">
          <h4 className="font-semibold mb-2"><code>ghotona_ekbar(emitter, event_name, callback)</code></h4>
          <p className="text-sm text-muted-foreground mb-2">
            Adds an event listener that runs only once, then is automatically removed.
          </p>
          <p className="text-sm"><strong>Returns:</strong> EventEmitter (for chaining)</p>
        </div>

        <div className="border rounded-lg p-4">
          <h4 className="font-semibold mb-2"><code>ghotona_prokash(emitter, event_name, ...data)</code></h4>
          <p className="text-sm text-muted-foreground mb-2">
            Emits an event with optional data arguments. All listeners for this event will be called.
          </p>
          <p className="text-sm"><strong>Returns:</strong> Boolean (sotti/true)</p>
        </div>

        <div className="border rounded-lg p-4">
          <h4 className="font-semibold mb-2"><code>ghotona_bondho(emitter, event_name, callback)</code></h4>
          <p className="text-sm text-muted-foreground mb-2">
            Removes a specific listener from an event.
          </p>
          <p className="text-sm"><strong>Returns:</strong> EventEmitter (for chaining)</p>
        </div>

        <div className="border rounded-lg p-4">
          <h4 className="font-semibold mb-2"><code>ghotona_sob_bondho(emitter, [event_name])</code></h4>
          <p className="text-sm text-muted-foreground mb-2">
            Removes all listeners for a specific event, or all listeners for all events if event_name is not provided.
          </p>
          <p className="text-sm"><strong>Returns:</strong> EventEmitter (for chaining)</p>
        </div>

        <div className="border rounded-lg p-4">
          <h4 className="font-semibold mb-2"><code>ghotona_shrotara(emitter, event_name)</code></h4>
          <p className="text-sm text-muted-foreground mb-2">
            Returns an array of all listeners for a specific event.
          </p>
          <p className="text-sm"><strong>Returns:</strong> Array of functions</p>
        </div>

        <div className="border rounded-lg p-4">
          <h4 className="font-semibold mb-2"><code>ghotona_naam_sob(emitter)</code></h4>
          <p className="text-sm text-muted-foreground mb-2">
            Returns an array of all event names that have listeners.
          </p>
          <p className="text-sm"><strong>Returns:</strong> Array of strings</p>
        </div>
      </div>

      <h2>Best Practices</h2>

      <div className="space-y-4 my-6">
        <div className="bg-green-50 dark:bg-green-950 border border-green-200 dark:border-green-800 rounded-lg p-4">
          <h4 className="text-green-900 dark:text-green-100 font-semibold mb-2">✅ Do</h4>
          <ul className="text-green-800 dark:text-green-200 space-y-1">
            <li>Use descriptive event names (e.g., "user:login", "data:received")</li>
            <li>Store listener references if you plan to remove them later</li>
            <li>Use <code>ghotona_ekbar()</code> for one-time events</li>
            <li>Document what events your components emit</li>
            <li>Clean up listeners when they're no longer needed</li>
          </ul>
        </div>

        <div className="bg-red-50 dark:bg-red-950 border border-red-200 dark:border-red-800 rounded-lg p-4">
          <h4 className="text-red-900 dark:text-red-100 font-semibold mb-2">❌ Don't</h4>
          <ul className="text-red-800 dark:text-red-200 space-y-1">
            <li>Don't create too many EventEmitters - reuse when possible</li>
            <li>Don't forget to remove listeners to prevent memory leaks</li>
            <li>Don't rely on listener execution order for critical logic</li>
            <li>Don't throw errors in listeners without handling them</li>
          </ul>
        </div>
      </div>

      <DocNavigation />
    </div>
  );
}
