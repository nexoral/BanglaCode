import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Path Utilities - BanglaCode Documentation",
  description: "Cross-platform path manipulation and file system utilities in BanglaCode",
};

export default function PathUtilitiesPage() {
  return (
    <div className="space-y-8">
      <section>
        <h1 className="text-4xl font-bold mb-4">Path Utilities (পাথ ইউটিলিটি)</h1>
        <p className="text-lg text-muted-foreground">
          BanglaCode provides comprehensive cross-platform path manipulation utilities. All operations are OS-aware and handle Windows, Linux, and macOS path conventions automatically.
        </p>
      </section>

      <section className="space-y-4">
        <h2 className="text-3xl font-semibold">Path Constants</h2>
        
        <div className="p-4 border rounded-lg">
          <h3 className="text-xl font-semibold mb-2">
            <code className="text-purple-600">PATH_SEP</code> &amp; <code className="text-purple-600">PATH_DELIMITER</code>
          </h3>
          <p className="text-muted-foreground mb-3">
            Platform-specific path separator and delimiter constants.
          </p>
          <div className="bg-slate-50 p-4 rounded">
            <pre className="text-sm overflow-x-auto">
              <code>{`dhoro sep = PATH_SEP;        // Unix: "/" | Windows: "\\\\"
dhoro delim = PATH_DELIMITER; // Unix: ":" | Windows: ";"`}</code>
            </pre>
          </div>
        </div>
      </section>

      <section className="space-y-4">
        <h2 className="text-3xl font-semibold">Path Functions</h2>

        <div className="space-y-6">
          <div className="p-4 border rounded-lg">
            <h3 className="text-xl font-semibold mb-2">
              <code className="text-blue-600">path_resolve(...paths)</code>
            </h3>
            <p className="text-muted-foreground mb-3">
              Resolves paths to absolute path.
            </p>
            <div className="bg-slate-50 p-4 rounded">
              <pre className="text-sm overflow-x-auto">
                <code>{`dhoro abs = path_resolve(".", "src", "main.bang");
dekho(abs);  // /home/user/project/src/main.bang`}</code>
              </pre>
            </div>
          </div>

          <div className="p-4 border rounded-lg">
            <h3 className="text-xl font-semibold mb-2">
              <code className="text-blue-600">path_normalize(path)</code>
            </h3>
            <p className="text-muted-foreground mb-3">
              Normalizes path by removing redundant separators and resolving . and ..
            </p>
            <div className="bg-slate-50 p-4 rounded">
              <pre className="text-sm overflow-x-auto">
                <code>{`dhoro clean = path_normalize("/foo//bar/../baz");
dekho(clean);  // /foo/baz`}</code>
              </pre>
            </div>
          </div>

          <div className="p-4 border rounded-lg">
            <h3 className="text-xl font-semibold mb-2">
              <code className="text-blue-600">path_relative(base, target)</code>
            </h3>
            <p className="text-muted-foreground mb-3">
              Computes relative path from base to target.
            </p>
            <div className="bg-slate-50 p-4 rounded">
              <pre className="text-sm overflow-x-auto">
                <code>{`dhoro rel = path_relative("/home/user", "/home/user/docs/file.txt");
dekho(rel);  // docs/file.txt`}</code>
              </pre>
            </div>
          </div>

          <div className="p-4 border rounded-lg">
            <h3 className="text-xl font-semibold mb-2">
              <code className="text-blue-600">path_joro(...segments)</code>
            </h3>
            <p className="text-muted-foreground mb-3">
              Joins path segments using platform-specific separator.
            </p>
            <div className="bg-slate-50 p-4 rounded">
              <pre className="text-sm overflow-x-auto">
                <code>{`dhoro fullPath = path_joro("home", "user", "file.txt");
// Unix: home/user/file.txt | Windows: home\\user\\file.txt`}</code>
              </pre>
            </div>
          </div>

          <div className="p-4 border rounded-lg">
            <h3 className="text-xl font-semibold mb-2">
              <code className="text-blue-600">directory_naam(path)</code>
            </h3>
            <p className="text-muted-foreground mb-3">
              Returns directory name of a path.
            </p>
            <div className="bg-slate-50 p-4 rounded">
              <pre className="text-sm overflow-x-auto">
                <code>{`dhoro dir = directory_naam("/home/user/file.txt");
dekho(dir);  // /home/user`}</code>
              </pre>
            </div>
          </div>

          <div className="p-4 border rounded-lg">
            <h3 className="text-xl font-semibold mb-2">
              <code className="text-blue-600">path_naam(path)</code>
            </h3>
            <p className="text-muted-foreground mb-3">
              Returns base name (filename or directory name).
            </p>
            <div className="bg-slate-50 p-4 rounded">
              <pre className="text-sm overflow-x-auto">
                <code>{`dhoro name = path_naam("/home/user/document.txt");
dekho(name);  // document.txt`}</code>
              </pre>
            </div>
          </div>

          <div className="p-4 border rounded-lg">
            <h3 className="text-xl font-semibold mb-2">
              <code className="text-blue-600">file_ext(path)</code>
            </h3>
            <p className="text-muted-foreground mb-3">
              Returns file extension (including dot).
            </p>
            <div className="bg-slate-50 p-4 rounded">
              <pre className="text-sm overflow-x-auto">
                <code>{`dhoro ext = file_ext("program.bang");
dekho(ext);  // .bang`}</code>
              </pre>
            </div>
          </div>
        </div>
      </section>

      <section className="space-y-4">
        <h2 className="text-3xl font-semibold">Real-World Example</h2>

        <div className="p-4 border rounded-lg bg-slate-50">
          <h3 className="text-xl font-semibold mb-3">Path Analysis Utility</h3>
          <pre className="text-sm overflow-x-auto">
            <code>{`kaj analyzePath(filePath) {
  ferao {
    absolute: path_resolve(filePath),
    normalized: path_normalize(filePath),
    directory: directory_naam(filePath),
    filename: path_naam(filePath),
    extension: file_ext(filePath),
    separator: PATH_SEP
  };
}

dhoro result = analyzePath("./src/../lib/utils.bang");
dekho(result);
// {
//   absolute: "/home/user/project/lib/utils.bang",
//   normalized: "lib/utils.bang",
//   directory: ".",
//   filename: "utils.bang",
//   extension: ".bang",
//   separator: "/"
// }`}</code>
          </pre>
        </div>
      </section>

      <section className="space-y-4">
        <h2 className="text-3xl font-semibold">API Summary</h2>
        <div className="overflow-x-auto">
          <table className="w-full border-collapse">
            <thead>
              <tr className="bg-slate-100">
                <th className="border p-2 text-left">Function/Constant</th>
                <th className="border p-2 text-left">Description</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td className="border p-2"><code>PATH_SEP</code></td>
                <td className="border p-2">Platform path separator</td>
              </tr>
              <tr>
                <td className="border p-2"><code>PATH_DELIMITER</code></td>
                <td className="border p-2">Platform path list delimiter</td>
              </tr>
              <tr>
                <td className="border p-2"><code>path_resolve(...paths)</code></td>
                <td className="border p-2">Resolve to absolute path</td>
              </tr>
              <tr>
                <td className="border p-2"><code>path_normalize(path)</code></td>
                <td className="border p-2">Normalize path</td>
              </tr>
              <tr>
                <td className="border p-2"><code>path_relative(base, target)</code></td>
                <td className="border p-2">Get relative path</td>
              </tr>
              <tr>
                <td className="border p-2"><code>path_joro(...segments)</code></td>
                <td className="border p-2">Join path segments</td>
              </tr>
              <tr>
                <td className="border p-2"><code>directory_naam(path)</code></td>
                <td className="border p-2">Get directory name</td>
              </tr>
              <tr>
                <td className="border p-2"><code>path_naam(path)</code></td>
                <td className="border p-2">Get base name</td>
              </tr>
              <tr>
                <td className="border p-2"><code>file_ext(path)</code></td>
                <td className="border p-2">Get file extension</td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>
    </div>
  );
}
