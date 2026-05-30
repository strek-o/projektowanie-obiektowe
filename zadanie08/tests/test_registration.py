import pytest
from selenium.webdriver.common.by import By
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.support.ui import WebDriverWait

VALID = {"username": "alice", "email": "alice@example.com", "password": "secret123"}


def _fill(page, **values):
    for field, value in values.items():
        element = page.find_element(By.ID, field)
        element.clear()
        element.send_keys(value)


def _submit(page):
    page.find_element(By.ID, "submit").click()


def _error(page, field):
    matches = page.find_elements(By.CSS_SELECTOR, f'[data-testid="error-{field}"]')
    return matches[0].text if matches else ""


def _wait_for_error(page, field, timeout=5):
    WebDriverWait(page, timeout).until(
        EC.visibility_of_element_located(
            (By.CSS_SELECTOR, f'[data-testid="error-{field}"]')
        )
    )


def _success_shown(page):
    return bool(page.find_elements(By.CSS_SELECTOR, '[data-testid="register-success"]'))


def test_registration_form_is_visible(page):
    assert page.find_element(By.ID, "username").is_displayed()
    assert page.find_element(By.ID, "email").is_displayed()
    assert page.find_element(By.ID, "password").is_displayed()
    assert page.find_element(By.ID, "submit").is_displayed()


def test_empty_form_shows_all_required_errors(page):
    _submit(page)

    _wait_for_error(page, "username")
    assert "required" in _error(page, "username").lower()
    assert "required" in _error(page, "email").lower()
    assert "required" in _error(page, "password").lower()
    assert not _success_shown(page), "empty form must not register a user"


@pytest.mark.parametrize("missing", ["username", "email", "password"])
def test_missing_required_field_is_reported(page, missing):
    values = {k: v for k, v in VALID.items() if k != missing}
    _fill(page, **values)
    _submit(page)

    _wait_for_error(page, missing)
    assert "required" in _error(page, missing).lower()
    for other in VALID:
        if other != missing:
            assert _error(page, other) == ""
    assert not _success_shown(page)


@pytest.mark.parametrize(
    "bad_email",
    [
        "plainaddress",
        "alice@",
        "@example.com",
        "alice@example",
        "alice example.com",
        "alice@@example.com",
    ],
)
def test_invalid_email_format_is_rejected(page, bad_email):
    _fill(page, username=VALID["username"], email=bad_email, password=VALID["password"])
    _submit(page)

    _wait_for_error(page, "email")
    assert "invalid email" in _error(page, "email").lower()
    assert not _success_shown(page), f"{bad_email!r} should not be accepted"


def test_invalid_email_does_not_affect_other_fields(page):
    _fill(page, username=VALID["username"], email="not-an-email", password=VALID["password"])
    _submit(page)

    _wait_for_error(page, "email")
    assert _error(page, "username") == ""
    assert _error(page, "password") == ""


def test_valid_data_registers_successfully(page):
    _fill(page, **VALID)
    _submit(page)

    WebDriverWait(page, 5).until(
        EC.visibility_of_element_located(
            (By.CSS_SELECTOR, '[data-testid="register-success"]')
        )
    )
    assert _success_shown(page)
    for field in VALID:
        assert _error(page, field) == ""


def test_correcting_email_then_registering_succeeds(page):
    _fill(page, username=VALID["username"], email="broken", password=VALID["password"])
    _submit(page)
    _wait_for_error(page, "email")
    assert not _success_shown(page)

    _fill(page, email=VALID["email"])
    _submit(page)
    WebDriverWait(page, 5).until(
        EC.visibility_of_element_located(
            (By.CSS_SELECTOR, '[data-testid="register-success"]')
        )
    )
    assert _error(page, "email") == ""
