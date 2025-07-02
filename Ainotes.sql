CREATE TABLE users (
  id uuid PRIMARY KEY,
  google_id varchar(255) UNIQUE NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  full_name varchar(255),
  avatar_url text,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE resources (
  id uuid PRIMARY KEY,
  user_id uuid,
  type varchar(255),
  source_url text,
  original_title varchar(255),
  status varchar(255),
  created_at timestamptz,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE notes (
  id uuid PRIMARY KEY,
  resource_id uuid UNIQUE,
  user_id uuid,
  title text,
  summary text,
  full_text text,
  created_at timestamptz,
  updated_at timestamptz,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE quizzes (
  id uuid PRIMARY KEY,
  note_id uuid,
  question text,
  options jsonb,
  correct_answer_index int,
  explanation text,
  FOREIGN KEY (note_id) REFERENCES notes(id)
);

CREATE TABLE flashcards (
  id uuid PRIMARY KEY,
  note_id uuid,
  front_text text,
  back_text text,
  FOREIGN KEY (note_id) REFERENCES notes(id)
);

CREATE TABLE resources_notes (
  resources_id uuid,
  notes_resource_id uuid,
  PRIMARY KEY (resources_id, notes_resource_id),
  FOREIGN KEY (resources_id) REFERENCES resources(id),
  FOREIGN KEY (notes_resource_id) REFERENCES notes(resource_id)
);
