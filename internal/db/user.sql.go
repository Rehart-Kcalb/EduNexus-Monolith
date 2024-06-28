// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const checkEnrollment = `-- name: CheckEnrollment :one
select enrolled_on from enrollments where course_id = $1 and user_id = $2
`

type CheckEnrollmentParams struct {
	CourseID int64 `json:"course_id"`
	UserID   int64 `json:"user_id"`
}

func (q *Queries) CheckEnrollment(ctx context.Context, arg CheckEnrollmentParams) (pgtype.Date, error) {
	row := q.db.QueryRow(ctx, checkEnrollment, arg.CourseID, arg.UserID)
	var enrolled_on pgtype.Date
	err := row.Scan(&enrolled_on)
	return enrolled_on, err
}

const createSubmission = `-- name: CreateSubmission :exec
Insert into submissions(content,assignment_id,info,user_id) values ($1,$2,$3,$4)
`

type CreateSubmissionParams struct {
	Content      []byte      `json:"content"`
	AssignmentID int64       `json:"assignment_id"`
	Info         pgtype.Text `json:"info"`
	UserID       int64       `json:"user_id"`
}

func (q *Queries) CreateSubmission(ctx context.Context, arg CreateSubmissionParams) error {
	_, err := q.db.Exec(ctx, createSubmission,
		arg.Content,
		arg.AssignmentID,
		arg.Info,
		arg.UserID,
	)
	return err
}

const dropCourse = `-- name: DropCourse :exec
Delete from enrollments where user_id = $1 and course_id = $2
`

type DropCourseParams struct {
	UserID   int64 `json:"user_id"`
	CourseID int64 `json:"course_id"`
}

func (q *Queries) DropCourse(ctx context.Context, arg DropCourseParams) error {
	_, err := q.db.Exec(ctx, dropCourse, arg.UserID, arg.CourseID)
	return err
}

const enrollIntoCourse = `-- name: EnrollIntoCourse :exec
INSERT INTO
  enrollments (course_id, user_id, enrolled_on)
VALUES
  (
    $1,
    $2,
    NOW()
  )
`

type EnrollIntoCourseParams struct {
	CourseID int64 `json:"course_id"`
	UserID   int64 `json:"user_id"`
}

func (q *Queries) EnrollIntoCourse(ctx context.Context, arg EnrollIntoCourseParams) error {
	_, err := q.db.Exec(ctx, enrollIntoCourse, arg.CourseID, arg.UserID)
	return err
}

const getLastGradesByCourse = `-- name: GetLastGradesByCourse :many
SELECT distinct(assignments.title), info, submissions.submitted_at FROM public.submissions 
	left join assignments on assignments.id = submissions.assignment_id
	left join courses on courses.id = assignments.course_id
where user_id = $1 and courses.title = $2
order by submissions.submitted_at desc
`

type GetLastGradesByCourseParams struct {
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
}

type GetLastGradesByCourseRow struct {
	Title       pgtype.Text `json:"title"`
	Info        pgtype.Text `json:"info"`
	SubmittedAt pgtype.Date `json:"submitted_at"`
}

func (q *Queries) GetLastGradesByCourse(ctx context.Context, arg GetLastGradesByCourseParams) ([]GetLastGradesByCourseRow, error) {
	rows, err := q.db.Query(ctx, getLastGradesByCourse, arg.UserID, arg.Title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetLastGradesByCourseRow{}
	for rows.Next() {
		var i GetLastGradesByCourseRow
		if err := rows.Scan(&i.Title, &i.Info, &i.SubmittedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMyCourses = `-- name: GetMyCourses :many
SELECT
  courses.title,
  courses.image,
  users.firstname AS organization_name,
  users.profile as organization_logo
FROM
  courses
  LEFT JOIN users ON users.id = courses.course_provider
WHERE
  users.id = courses.course_provider
  AND courses.id IN (
    SELECT
      courses.id
    FROM
      courses
      INNER JOIN enrollments ON enrollments.course_id = courses.id
    WHERE
      enrollments.user_id = $1
  )
`

type GetMyCoursesRow struct {
	Title            string      `json:"title"`
	Image            pgtype.Text `json:"image"`
	OrganizationName pgtype.Text `json:"organization_name"`
	OrganizationLogo pgtype.Text `json:"organization_logo"`
}

func (q *Queries) GetMyCourses(ctx context.Context, userID int64) ([]GetMyCoursesRow, error) {
	rows, err := q.db.Query(ctx, getMyCourses, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetMyCoursesRow{}
	for rows.Next() {
		var i GetMyCoursesRow
		if err := rows.Scan(
			&i.Title,
			&i.Image,
			&i.OrganizationName,
			&i.OrganizationLogo,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProfileInfo = `-- name: GetProfileInfo :one
select firstname,description,profile from users where users.id = $1
`

type GetProfileInfoRow struct {
	Firstname   pgtype.Text `json:"firstname"`
	Description pgtype.Text `json:"description"`
	Profile     pgtype.Text `json:"profile"`
}

func (q *Queries) GetProfileInfo(ctx context.Context, id int64) (GetProfileInfoRow, error) {
	row := q.db.QueryRow(ctx, getProfileInfo, id)
	var i GetProfileInfoRow
	err := row.Scan(&i.Firstname, &i.Description, &i.Profile)
	return i, err
}

const getReadedLecturesByModule = `-- name: GetReadedLecturesByModule :many
SELECT 
    distinct(m.id) AS module_id,
    m.title AS module_name,
    a.title,
    a.id AS assignment_id,
	a.assignment_type_id,
    COALESCE(pr.done IS NOT NULL, FALSE) AS read
  FROM 
    modules m
  LEFT JOIN 
    assignments a ON m.id = a.module_id
  LEFT JOIN 
    progress pr ON a.id = pr.assignment_id
where m.id = $1 and pr.user_id = $2
`

type GetReadedLecturesByModuleParams struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

type GetReadedLecturesByModuleRow struct {
	ModuleID         int64       `json:"module_id"`
	ModuleName       string      `json:"module_name"`
	Title            pgtype.Text `json:"title"`
	AssignmentID     pgtype.Int8 `json:"assignment_id"`
	AssignmentTypeID pgtype.Int8 `json:"assignment_type_id"`
	Read             interface{} `json:"read"`
}

func (q *Queries) GetReadedLecturesByModule(ctx context.Context, arg GetReadedLecturesByModuleParams) ([]GetReadedLecturesByModuleRow, error) {
	rows, err := q.db.Query(ctx, getReadedLecturesByModule, arg.ID, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetReadedLecturesByModuleRow{}
	for rows.Next() {
		var i GetReadedLecturesByModuleRow
		if err := rows.Scan(
			&i.ModuleID,
			&i.ModuleName,
			&i.Title,
			&i.AssignmentID,
			&i.AssignmentTypeID,
			&i.Read,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProfileInfo = `-- name: UpdateProfileInfo :exec
Update users set firstname = $1, description = $2, profile = $3 where users.id = $4
`

type UpdateProfileInfoParams struct {
	Firstname   pgtype.Text `json:"firstname"`
	Description pgtype.Text `json:"description"`
	Profile     pgtype.Text `json:"profile"`
	ID          int64       `json:"id"`
}

func (q *Queries) UpdateProfileInfo(ctx context.Context, arg UpdateProfileInfoParams) error {
	_, err := q.db.Exec(ctx, updateProfileInfo,
		arg.Firstname,
		arg.Description,
		arg.Profile,
		arg.ID,
	)
	return err
}