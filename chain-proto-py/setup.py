from setuptools import find_namespace_packages, setup

with open('requirements.txt') as f:
    required = f.read().splitlines()

setup(
    name="chain-proto",
    version="0.0.0",
    author="Bitoro Trading Inc.",
    author_email="contact@bitoro.foundation",
    description="Protos for Bitoro Chain protocol",
    packages = find_namespace_packages(),
    include_package_data=True,  # Include files specified in MANIFEST.in
    install_requires=required,
    license_files = ("LICENSE"),
    python_requires=">=3.8",
)
