import { useState } from 'react';
import { Container, Row, Col, Form, Button, Alert } from 'react-bootstrap';
import './App.css';

function App() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [feedback, setFeedback] = useState('');
  const [submitting, setSubmitting] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const handleSubmit = (event) => {
    event.preventDefault();
    setSubmitting(true);
    setSuccess(false);
    setError(false);

    fetch('http://localhost:8080/feedback/create', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name, email, feedback }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        setSuccess(true);
      })
      .catch((error) => {
        console.error(error);
        setError(true);
      })
      .finally(() => {
        setSubmitting(false);
        setName('');
        setEmail('');
        setFeedback('');
      });
  };

  return (
    <Container>
      <Row>
        <Col md={{ span: 6, offset: 3 }}>
          <h2>Submit Your Feedback</h2>
          {success && <Alert variant="success">Feedback submitted successfully!</Alert>}
          {error && <Alert variant="danger">Failed to submit feedback. Please try again.</Alert>}
          <Form onSubmit={handleSubmit}>
            <Form.Group controlId="name">
              <Form.Label>Name</Form.Label>
              <Form.Control
                type="text"
                value={name}
                onChange={(event) => setName(event.target.value)}
                placeholder="Enter your name"
              />
            </Form.Group>
            <Form.Group controlId="email">
              <Form.Label>Email Address</Form.Label>
              <Form.Control
                type="email"
                value={email}
                onChange={(event) => setEmail(event.target.value)}
                placeholder="Enter your email address"
              />
            </Form.Group>
            <Form.Group controlId="feedback">
              <Form.Label>Feedback</Form.Label>
              <Form.Control
                as="textarea"
                rows={5}
                value={feedback}
                onChange={(event) => setFeedback(event.target.value)}
                placeholder="Enter your feedback"
              />
            </Form.Group>
            <Button variant="primary" type="submit" disabled={submitting}>
              {submitting ? 'Submitting...' : 'Submit'}
            </Button>
          </Form>
        </Col>
      </Row>
    </Container>
  );
}

export default App;