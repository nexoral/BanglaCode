export const metadata = {
  title: "Collections (Set & Map) - BanglaCode Documentation",
  description:
    "Learn Set and Map collections in BanglaCode for unique values and flexible key-value storage.",
};

export default function CollectionsPage() {
  return (
    <div className="max-w-4xl mx-auto px-6 py-10 space-y-8">
      <div>
        <h1 className="text-4xl font-bold mb-3">Collections: Set &amp; Map</h1>
        <p className="text-lg text-gray-700 dark:text-gray-300">
          BanglaCode supports ES6-style collections: <code>Set</code> for unique values and
          <code> Map</code> for key-value pairs with flexible keys.
        </p>
      </div>

      <section className="space-y-3">
        <h2 className="text-2xl font-semibold">Quick Start</h2>
        <div className="bg-gray-100 dark:bg-gray-900 rounded-lg p-4">
          <pre className="text-sm overflow-x-auto">
            <code className="language-banglacode">
{`// Set - unique values
dhoro mySet = set_srishti([1, 2, 3, 2, 1]);
dekho(set_akar(mySet)); // 3

// Map - key-value pairs
dhoro myMap = map_srishti();
map_set(myMap, "name", "Ankan");
map_set(myMap, [1, 2], "array key");
dekho(map_get(myMap, "name")); // "Ankan"`}
            </code>
          </pre>
        </div>
      </section>

      <section className="space-y-3">
        <h2 className="text-2xl font-semibold">Set API</h2>
        <ul className="list-disc list-inside space-y-1">
          <li><code>set_srishti(initialArray?)</code> - Create a new Set.</li>
          <li><code>set_add(set, value)</code> - Add unique value.</li>
          <li><code>set_has(set, value)</code> - Membership check.</li>
          <li><code>set_delete(set, value)</code> - Remove value.</li>
          <li><code>set_clear(set)</code> - Remove all values.</li>
          <li><code>set_akar(set)</code> - Size of set.</li>
          <li><code>set_values(set)</code> - Convert set to array.</li>
          <li><code>set_foreach(set, callback)</code> - Iterate values.</li>
        </ul>
      </section>

      <section className="space-y-3">
        <h2 className="text-2xl font-semibold">Map API</h2>
        <ul className="list-disc list-inside space-y-1">
          <li><code>map_srishti(entries?)</code> - Create a new Map.</li>
          <li><code>map_set(map, key, value)</code> - Insert/update entry.</li>
          <li><code>map_get(map, key)</code> - Get value by key.</li>
          <li><code>map_has(map, key)</code> - Check key existence.</li>
          <li><code>map_delete(map, key)</code> - Delete key.</li>
          <li><code>map_clear(map)</code> - Clear all entries.</li>
          <li><code>map_akar(map)</code> - Number of entries.</li>
          <li><code>map_keys(map)</code>, <code>map_values(map)</code>, <code>map_entries(map)</code>.</li>
          <li><code>map_foreach(map, callback)</code> - Iterate entries.</li>
        </ul>
      </section>

      <section className="space-y-3">
        <h2 className="text-2xl font-semibold">Example: De-duplicate + Count</h2>
        <div className="bg-gray-100 dark:bg-gray-900 rounded-lg p-4">
          <pre className="text-sm overflow-x-auto">
            <code className="language-banglacode">
{`dhoro words = ["a", "b", "a", "c", "b", "a"];

// Unique words
dhoro uniqueWords = set_values(set_srishti(words));
dekho(uniqueWords); // ["a", "b", "c"]

// Frequency count
dhoro freq = map_srishti();
ghuriye (dhoro i = 0; i < dorghyo(words); i = i + 1) {
  dhoro w = words[i];
  dhoro n = map_get(freq, w);
  jodi (n == khali) {
    map_set(freq, w, 1);
  } nahole {
    map_set(freq, w, n + 1);
  }
}

map_foreach(freq, kaj(value, key) {
  dekho(key, "=>", value);
});`}
            </code>
          </pre>
        </div>
      </section>
    </div>
  );
}

