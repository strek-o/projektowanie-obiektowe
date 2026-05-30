import functools
import http.server
import os
import socketserver
import threading
from pathlib import Path

import pytest
from selenium import webdriver
from selenium.webdriver.chrome.options import Options

APP_DIR = Path(__file__).resolve().parent.parent / "app"


@pytest.fixture(scope="session")
def base_url():
    override = os.environ.get("BASE_URL")
    if override:
        yield override.rstrip("/")
        return

    handler = functools.partial(
        http.server.SimpleHTTPRequestHandler, directory=str(APP_DIR)
    )
    server = socketserver.TCPServer(("127.0.0.1", 0), handler)
    thread = threading.Thread(target=server.serve_forever, daemon=True)
    thread.start()

    host, port = server.server_address
    try:
        yield f"http://{host}:{port}"
    finally:
        server.shutdown()
        server.server_close()


@pytest.fixture(scope="session")
def driver():
    options = Options()
    options.add_argument("--no-sandbox")
    options.add_argument("--disable-gpu")
    options.add_argument("--disable-dev-shm-usage")
    options.add_argument("--window-size=1280,1024")

    chrome = webdriver.Chrome(options=options)
    chrome.implicitly_wait(0)
    try:
        yield chrome
    finally:
        chrome.quit()


@pytest.fixture
def page(driver, base_url):
    driver.get(f"{base_url}/index.html")
    return driver
