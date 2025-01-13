import codeinfra
import codeinfra_synthetic as synthetic

rt = synthetic.resource_properties.Root("rt")
codeinfra.export("trivial", rt)
codeinfra.export("simple", rt.res1)
codeinfra.export("foo", rt.res1.obj1.res2.obj2)
codeinfra.export("complex", rt.res1.obj1.res2.obj2.answer)
