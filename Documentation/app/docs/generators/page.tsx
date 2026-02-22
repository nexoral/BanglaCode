export const metadata = {
  title: "Generators - BanglaCode",
  description: "Learn generator functions with kaj* and utpadan in BanglaCode.",
};

export default function GeneratorsDoc() {
  return (
    <div className="max-w-4xl mx-auto px-6 py-10">
      <h1 className="text-4xl font-bold mb-6">Generators</h1>

      <p className="text-lg text-gray-700 dark:text-gray-300 mb-8">
        BanglaCode generators let you produce values lazily using <code>kaj*</code> and{" "}
        <code>utpadan</code>. They return a generator object with <code>next()</code>,{" "}
        <code>return()</code>, and <code>throw()</code>.
      </p>

      <section className="mb-10">
        <h2 className="text-2xl font-bold mb-4">Basic Generator</h2>
        <div className="bg-gray-100 dark:bg-gray-900 rounded-lg p-5">
          <pre className="text-sm overflow-x-auto">
            <code className="language-banglacode">
{`kaj* count(max) {
  dhoro i = 0;
  jotokkhon (i < max) {
    utpadan i;
    i = i + 1;
  }
}

dhoro g = count(3);
dekho(g.next()); // {"value": 0, "done": mittha}
dekho(g.next()); // {"value": 1, "done": mittha}
dekho(g.next()); // {"value": 2, "done": mittha}
dekho(g.next()); // {"value": khali, "done": sotti}`}
            </code>
          </pre>
        </div>
      </section>

      <section className="mb-10">
        <h2 className="text-2xl font-bold mb-4">Generator Methods</h2>
        <ul className="list-disc list-inside space-y-2 text-gray-700 dark:text-gray-300">
          <li><code>next()</code>: Resume execution until next <code>utpadan</code> or completion.</li>
          <li><code>return(value)</code>: Stop generator early and mark done.</li>
          <li><code>throw(err)</code>: Throw an error into the generator and end it.</li>
        </ul>
      </section>

      <section>
        <h2 className="text-2xl font-bold mb-4">Early Return Example</h2>
        <div className="bg-gray-100 dark:bg-gray-900 rounded-lg p-5">
          <pre className="text-sm overflow-x-auto">
            <code className="language-banglacode">
{`kaj* ids() {
  utpadan 101;
  utpadan 102;
  utpadan 103;
}

dhoro g = ids();
dekho(g.next()["value"]);      // 101
dekho(g.return("stop"));       // {"value": "stop", "done": sotti}
dekho(g.next()["done"]);       // sotti`}
            </code>
          </pre>
        </div>
      </section>
    </div>
  );
}

