package persistence

var _ DataBaseInterface = (*DataBase)(nil)

//type PersistenceTestSuite struct {
//	suite.Suite
//	db       DataBaseInterface
//	original DataBaseInterface
//}
//
//func (s *PersistenceTestSuite) SetupTest() {
//	s.db = &MockDb{}
//	s.original = Db
//	Db = s.db
//}
//
//
//func (s *PersistenceTestSuite) TearDownTest() {
//	Db = s.original
//}
//
//func TestPersistenceSuite(t *testing.T) {
//	suite.Run(t, new(PersistenceTestSuite))
//}
