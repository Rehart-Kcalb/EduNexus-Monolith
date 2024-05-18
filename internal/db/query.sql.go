// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const allCategories = `-- name: AllCategories :many
SELECT
  "name",
  "color"
FROM
  categories
`

type AllCategoriesRow struct {
	Name  string `json:"name"`
	Color int32  `json:"color"`
}

func (q *Queries) AllCategories(ctx context.Context) ([]AllCategoriesRow, error) {
	rows, err := q.db.Query(ctx, allCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AllCategoriesRow
	for rows.Next() {
		var i AllCategoriesRow
		if err := rows.Scan(&i.Name, &i.Color); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

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

const countCourses = `-- name: CountCourses :one
select count(title) from filter($1,$2::bigint[])
`

type CountCoursesParams struct {
	TitleParam string  `json:"title_param"`
	Column2    []int64 `json:"column_2"`
}

func (q *Queries) CountCourses(ctx context.Context, arg CountCoursesParams) (int64, error) {
	row := q.db.QueryRow(ctx, countCourses, arg.TitleParam, arg.Column2)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO
  users ("login", "password", "user_role_id")
VALUES
  ($1, $2, 1)
`

type CreateUserParams struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser, arg.Login, arg.Password)
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

const filterCourses = `-- name: FilterCourses :many
select title,image,organization_name,organization_logo from filter($1,$2::bigint[]) limit $3 offset $4
`

type FilterCoursesParams struct {
	TitleParam string  `json:"title_param"`
	Column2    []int64 `json:"column_2"`
	Limit      int32   `json:"limit"`
	Offset     int32   `json:"offset"`
}

func (q *Queries) FilterCourses(ctx context.Context, arg FilterCoursesParams) ([]GetMyCoursesRow, error) {
	rows, err := q.db.Query(ctx, filterCourses,
		arg.TitleParam,
		arg.Column2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMyCoursesRow
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

const getCategoryId = `-- name: GetCategoryId :one
select id from categories where name=$1 limit 1
`

func (q *Queries) GetCategoryId(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRow(ctx, getCategoryId, name)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getClaimsByLogin = `-- name: GetClaimsByLogin :one
SELECT
  users.id,
  user_roles.title
FROM
  users
  LEFT JOIN user_roles ON user_roles.id = users.user_role_id
WHERE
  users.login = $1
`

type GetClaimsByLoginRow struct {
	ID    int64       `json:"id"`
	Title pgtype.Text `json:"title"`
}

func (q *Queries) GetClaimsByLogin(ctx context.Context, login string) (GetClaimsByLoginRow, error) {
	row := q.db.QueryRow(ctx, getClaimsByLogin, login)
	var i GetClaimsByLoginRow
	err := row.Scan(&i.ID, &i.Title)
	return i, err
}

const getCourseDetails = `-- name: GetCourseDetails :one
SELECT
  c.description,c.id,c.image
FROM
  courses c where c.id = $1
`

type GetCourseDetailsRow struct {
	Description string      `json:"description"`
	ID          int64       `json:"id"`
	Image       pgtype.Text `json:"image"`
}

func (q *Queries) GetCourseDetails(ctx context.Context, id int64) (GetCourseDetailsRow, error) {
	row := q.db.QueryRow(ctx, getCourseDetails, id)
	var i GetCourseDetailsRow
	err := row.Scan(&i.Description, &i.ID, &i.Image)
	return i, err
}

const getCourseEnrolledAmount = `-- name: GetCourseEnrolledAmount :one
select count(id) from enrollments where enrollments.course_id = $1
`

func (q *Queries) GetCourseEnrolledAmount(ctx context.Context, courseID int64) (int64, error) {
	row := q.db.QueryRow(ctx, getCourseEnrolledAmount, courseID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getCourseId = `-- name: GetCourseId :one
select id from courses where title = $1 limit 1
`

func (q *Queries) GetCourseId(ctx context.Context, title string) (int64, error) {
	row := q.db.QueryRow(ctx, getCourseId, title)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getCourseLectures = `-- name: GetCourseLectures :many
select a.id, a.module_id, a.course_id, a.title, a.description, a.content, a.days, a.assignment_type_id from courses c 
left join modules m on m.course_id = c.id
left join assignments a on a.module_id = m.id
where  c.id = $1 and a.id is not null
`

type GetCourseLecturesRow struct {
	ID               pgtype.Int8 `json:"id"`
	ModuleID         pgtype.Int8 `json:"module_id"`
	CourseID         pgtype.Int8 `json:"course_id"`
	Title            pgtype.Text `json:"title"`
	Description      pgtype.Text `json:"description"`
	Content          pgtype.Text `json:"content"`
	Days             pgtype.Int4 `json:"days"`
	AssignmentTypeID pgtype.Int8 `json:"assignment_type_id"`
}

func (q *Queries) GetCourseLectures(ctx context.Context, id int64) ([]GetCourseLecturesRow, error) {
	rows, err := q.db.Query(ctx, getCourseLectures, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCourseLecturesRow
	for rows.Next() {
		var i GetCourseLecturesRow
		if err := rows.Scan(
			&i.ID,
			&i.ModuleID,
			&i.CourseID,
			&i.Title,
			&i.Description,
			&i.Content,
			&i.Days,
			&i.AssignmentTypeID,
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

const getCourseModules = `-- name: GetCourseModules :many
SELECT
  modules.title
FROM
  modules
  INNER JOIN courses ON courses.id = modules.course_id
WHERE
  courses.title = $1
`

func (q *Queries) GetCourseModules(ctx context.Context, title string) ([]string, error) {
	rows, err := q.db.Query(ctx, getCourseModules, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			return nil, err
		}
		items = append(items, title)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCourseTeachers = `-- name: GetCourseTeachers :many
SELECT
  u.firstname,u.surname,u.profile
FROM
  courses
  INNER JOIN course_teachers ct ON ct.course_id = courses.id 
  inner join users u on u.id = ct.user_id
  where ct.course_id = $1 AND u.user_role_id = 1
`

type GetCourseTeachersRow struct {
	Firstname pgtype.Text `json:"firstname"`
	Surname   pgtype.Text `json:"surname"`
	Profile   pgtype.Text `json:"profile"`
}

func (q *Queries) GetCourseTeachers(ctx context.Context, courseID int64) ([]GetCourseTeachersRow, error) {
	rows, err := q.db.Query(ctx, getCourseTeachers, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCourseTeachersRow
	for rows.Next() {
		var i GetCourseTeachersRow
		if err := rows.Scan(&i.Firstname, &i.Surname, &i.Profile); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLectureContent = `-- name: GetLectureContent :one
select id, module_id, course_id, title, description, content, days, assignment_type_id from assignments
where assignments.id = $1
`

func (q *Queries) GetLectureContent(ctx context.Context, id int64) (Assignment, error) {
	row := q.db.QueryRow(ctx, getLectureContent, id)
	var i Assignment
	err := row.Scan(
		&i.ID,
		&i.ModuleID,
		&i.CourseID,
		&i.Title,
		&i.Description,
		&i.Content,
		&i.Days,
		&i.AssignmentTypeID,
	)
	return i, err
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
	var items []GetMyCoursesRow
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

const getPasswordByLogin = `-- name: GetPasswordByLogin :one
SELECT
  PASSWORD
FROM
  users
WHERE
  login = $1
`

func (q *Queries) GetPasswordByLogin(ctx context.Context, login string) (string, error) {
	row := q.db.QueryRow(ctx, getPasswordByLogin, login)
	var password string
	err := row.Scan(&password)
	return password, err
}

const newLecture = `-- name: NewLecture :exec
insert into assignments (module_id,title,description,content,assignment_type_id)
values ($1,$2,$3,$4,1)
`

type NewLectureParams struct {
	ModuleID    int64       `json:"module_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Content     pgtype.Text `json:"content"`
}

func (q *Queries) NewLecture(ctx context.Context, arg NewLectureParams) error {
	_, err := q.db.Exec(ctx, newLecture,
		arg.ModuleID,
		arg.Title,
		arg.Description,
		arg.Content,
	)
	return err
}
