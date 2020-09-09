
try:
    import importlib.resources as pkg_resources
except ImportError:
    # Try backported to PY<37 `importlib_resources`.
    import importlib_resources as pkg_resources

#from . import templates  # relative-import the *package* containing the templates

import jinja2

import collections
import os

class RenderError(Exception):
    pass

def fq_name(name):
    return "::".join(name)

def cpp_namespace(name):
    u = name.split("::")
    if len(u) <= 1:
        return ""
    return "::".join(u[:-1])

def generate_golang(schema):
    env = jinja2.Environment(
            loader=jinja2.PackageLoader(__package__, "templates"),
            keep_trailing_newline=True,
        )
    env.filters["fq_name"] = fq_name
    env.filters["tuple"] = tuple
    decls_by_name = collections.OrderedDict(((fq_name(name), decl) for name, decl in schema["decls"]))
    decl_namespaces = sorted(set(cpp_namespace(name) for name in decls_by_name))

    ctx = {"schema" : schema,
           "decls_by_name" : decls_by_name,
           "decl_namespaces" : decl_namespaces,
          }
    for name, val in ctx["decls_by_name"].items():
        print(name)
        import json
        print(json.dumps(val))

    result = collections.OrderedDict()
    result_files = collections.OrderedDict()
    result["files"] = result_files

    template_names = [
        "koinos.go.j2",
        ]

    for template_name in template_names:
        j2_template = env.get_template(template_name)
        out_filename = os.path.splitext(template_name)[0]
        result_files[out_filename] = j2_template.render(ctx)

    rt_path = os.path.join(os.path.dirname(__file__), "rt")
    for root, dirs, files in os.walk(rt_path):
        for f in files:
            filepath = os.path.join(root, f)
            relpath = os.path.relpath(filepath, rt_path)
            with open(filepath, "r") as f:
                content = f.read()
            result_files[os.path.join("rt", relpath)] = content
    return result

def setup(app):
    app.register_target("golang", generate_golang)