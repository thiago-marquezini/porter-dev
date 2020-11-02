package gorm_test

import (
	"testing"

	"github.com/porter-dev/porter/internal/repository"

	"github.com/go-test/deep"
	"github.com/porter-dev/porter/internal/models"
	orm "gorm.io/gorm"
)

func initServiceAccountCandidate(tester *tester, t *testing.T) {
	t.Helper()

	saCandidate := &models.ServiceAccountCandidate{
		ProjectID:       1,
		Kind:            "connector",
		ClusterName:     "cluster-test",
		ClusterEndpoint: "https://localhost",
		AuthMechanism:   models.X509,
		Kubeconfig:      []byte("current-context: testing\n"),
		Actions: []models.ServiceAccountAction{
			models.ServiceAccountAction{
				Name:     models.TokenDataAction,
				Resolved: false,
			},
		},
	}

	saCandidate, err := tester.repo.ServiceAccount.CreateServiceAccountCandidate(saCandidate)

	if err != nil {
		t.Fatalf("%v\n", err)
	}

	tester.initSACandidates = append(tester.initSACandidates, saCandidate)
}

func initServiceAccount(tester *tester, t *testing.T) {
	t.Helper()

	sa := &models.ServiceAccount{
		ProjectID:             1,
		Kind:                  "connector",
		AuthMechanism:         models.X509,
		ClientCertificateData: []byte("-----BEGIN"),
		ClientKeyData:         []byte("-----BEGIN"),
		Clusters: []models.Cluster{
			models.Cluster{
				Name:                     "cluster-test",
				Server:                   "https://localhost",
				CertificateAuthorityData: []byte("-----BEGIN"),
			},
		},
	}

	sa, err := tester.repo.ServiceAccount.CreateServiceAccount(sa)

	if err != nil {
		t.Fatalf("%v\n", err)
	}

	tester.initSAs = append(tester.initSAs, sa)
}

func TestCreateServiceAccountCandidate(t *testing.T) {
	tester := &tester{
		dbFileName: "./porter_create_sa_candidate.db",
	}

	setupTestEnv(tester, t)
	initProject(tester, t)
	defer cleanup(tester, t)

	saCandidate := &models.ServiceAccountCandidate{
		ProjectID:       1,
		Kind:            "connector",
		ClusterName:     "cluster-test",
		ClusterEndpoint: "https://localhost",
		AuthMechanism:   models.X509,
		Kubeconfig:      []byte("current-context: testing\n"),
	}

	saCandidate, err := tester.repo.ServiceAccount.CreateServiceAccountCandidate(saCandidate)

	if err != nil {
		t.Fatalf("%v\n", err)
	}

	saCandidate, err = tester.repo.ServiceAccount.ReadServiceAccountCandidate(saCandidate.Model.ID)

	if err != nil {
		t.Fatalf("%v\n", err)
	}

	// make sure id is 1
	if saCandidate.Model.ID != 1 {
		t.Errorf("incorrect service accound candidate ID: expected %d, got %d\n", 1, saCandidate.Model.ID)
	}

	// make sure data is correct
	expSACandidate := &models.ServiceAccountCandidate{
		ProjectID:       1,
		Kind:            "connector",
		ClusterName:     "cluster-test",
		ClusterEndpoint: "https://localhost",
		AuthMechanism:   models.X509,
		Kubeconfig:      []byte("current-context: testing\n"),
		Actions:         []models.ServiceAccountAction{},
	}

	copySACandidate := saCandidate

	// reset fields for reflect.DeepEqual
	copySACandidate.Model = orm.Model{}
	copySACandidate.Kubeconfig, _ = repository.Decrypt(
		copySACandidate.Kubeconfig,
		tester.key,
	)

	if diff := deep.Equal(copySACandidate, expSACandidate); diff != nil {
		t.Errorf("incorrect sa candidate")
		t.Error(diff)
	}
}

func TestCreateServiceAccountCandidateWithAction(t *testing.T) {
	tester := &tester{
		dbFileName: "./porter_create_sa_candidate_w_action.db",
	}

	setupTestEnv(tester, t)
	initProject(tester, t)
	initServiceAccountCandidate(tester, t)
	defer cleanup(tester, t)

	saCandidate := tester.initSACandidates[0]

	saCandidate, err := tester.repo.ServiceAccount.ReadServiceAccountCandidate(saCandidate.Model.ID)

	if err != nil {
		t.Fatalf("%v\n", err)
	}

	// make sure IDs are correct
	if saCandidate.Model.ID != 1 {
		t.Errorf("incorrect service account candidate ID: expected %d, got %d\n", 1, saCandidate.Model.ID)
	}

	if len(saCandidate.Actions) != 1 {
		t.Errorf("incorrect actions length: expected %d, got %d\n", 1, len(saCandidate.Actions))
	}

	if saCandidate.Actions[0].Model.ID != 1 {
		t.Errorf("incorrect actions ID: expected %d, got %d\n", 1, saCandidate.Actions[0].Model.ID)
	}

	// make sure data is correct
	expSACandidate := &models.ServiceAccountCandidate{
		ProjectID:       1,
		Kind:            "connector",
		ClusterName:     "cluster-test",
		ClusterEndpoint: "https://localhost",
		AuthMechanism:   models.X509,
		Kubeconfig:      []byte("current-context: testing\n"),
		Actions: []models.ServiceAccountAction{
			models.ServiceAccountAction{
				ServiceAccountCandidateID: 1,
				Name:                      models.TokenDataAction,
				Resolved:                  false,
			},
		},
	}

	copySACandidate := saCandidate

	// reset fields for reflect.DeepEqual
	copySACandidate.Model = orm.Model{}
	copySACandidate.Kubeconfig, _ = repository.Decrypt(
		copySACandidate.Kubeconfig,
		tester.key,
	)
	copySACandidate.Actions[0].Model = orm.Model{}

	if diff := deep.Equal(copySACandidate, expSACandidate); diff != nil {
		t.Errorf("incorrect sa candidate")
		t.Error(diff)
	}
}

func TestListServiceAccountCandidatesByProjectID(t *testing.T) {
	tester := &tester{
		dbFileName: "./porter_list_sa_candidates.db",
	}

	setupTestEnv(tester, t)
	initProject(tester, t)
	initServiceAccountCandidate(tester, t)
	defer cleanup(tester, t)

	saCandidates, err := tester.repo.ServiceAccount.ListServiceAccountCandidatesByProjectID(
		tester.initProjects[0].Model.ID,
	)

	if err != nil {
		t.Fatalf("%v\n", err)
	}

	if len(saCandidates) != 1 {
		t.Fatalf("length of sa candidates incorrect: expected %d, got %d\n", 1, len(saCandidates))
	}

	// make sure data is correct
	expSACandidate := &models.ServiceAccountCandidate{
		ProjectID:       1,
		Kind:            "connector",
		ClusterName:     "cluster-test",
		ClusterEndpoint: "https://localhost",
		AuthMechanism:   models.X509,
		Kubeconfig:      []byte("current-context: testing\n"),
		Actions: []models.ServiceAccountAction{
			models.ServiceAccountAction{
				ServiceAccountCandidateID: 1,
				Name:                      models.TokenDataAction,
				Resolved:                  false,
			},
		},
	}

	copySACandidate := saCandidates[0]

	// reset fields for reflect.DeepEqual
	copySACandidate.Model = orm.Model{}
	copySACandidate.Kubeconfig, _ = repository.Decrypt(
		copySACandidate.Kubeconfig,
		tester.key,
	)
	copySACandidate.Actions[0].Model = orm.Model{}

	if diff := deep.Equal(copySACandidate, expSACandidate); diff != nil {
		t.Errorf("incorrect sa candidate")
		t.Error(diff)
	}
}

func TestCreateServiceAccount(t *testing.T) {
	tester := &tester{
		dbFileName: "./porter_create_sa.db",
	}

	setupTestEnv(tester, t)
	initProject(tester, t)
	defer cleanup(tester, t)

	sa := &models.ServiceAccount{
		ProjectID:             1,
		Kind:                  "connector",
		AuthMechanism:         models.X509,
		ClientCertificateData: []byte("-----BEGIN"),
		ClientKeyData:         []byte("-----BEGIN"),
	}

	sa, err := tester.repo.ServiceAccount.CreateServiceAccount(sa)

	if err != nil {
		t.Fatalf("%v\n", err)
	}

	sa, err = tester.repo.ServiceAccount.ReadServiceAccount(sa.Model.ID)

	if err != nil {
		t.Fatalf("%v\n", err)
	}

	// make sure id is 1
	if sa.Model.ID != 1 {
		t.Errorf("incorrect service account ID: expected %d, got %d\n", 1, sa.Model.ID)
	}

	// make sure data is correct
	expSA := &models.ServiceAccount{
		ProjectID:             1,
		Kind:                  "connector",
		AuthMechanism:         models.X509,
		ClientCertificateData: []byte("-----BEGIN"),
		ClientKeyData:         []byte("-----BEGIN"),
		Clusters:              []models.Cluster{},
	}

	copySA := sa

	// reset fields for reflect.DeepEqual
	copySA.Model = orm.Model{}
	copySA.ClientCertificateData, _ = repository.Decrypt(copySA.ClientCertificateData, tester.key)
	copySA.ClientKeyData, _ = repository.Decrypt(copySA.ClientKeyData, tester.key)

	if diff := deep.Equal(copySA, expSA); diff != nil {
		t.Errorf("incorrect service account")
		t.Error(diff)
	}
}

func TestCreateServiceAccountWithCluster(t *testing.T) {
	tester := &tester{
		dbFileName: "./porter_create_sa_w_cluster.db",
	}

	setupTestEnv(tester, t)
	initProject(tester, t)
	initServiceAccount(tester, t)
	defer cleanup(tester, t)

	sa := tester.initSAs[0]

	sa, err := tester.repo.ServiceAccount.ReadServiceAccount(sa.Model.ID)

	if err != nil {
		t.Fatalf("%v\n", err)
	}

	// make sure id is 1
	if sa.Model.ID != 1 {
		t.Errorf("incorrect service account ID: expected %d, got %d\n", 1, sa.Model.ID)
	}

	if len(sa.Clusters) != 1 {
		t.Errorf("incorrect clusters length: expected %d, got %d\n", 1, len(sa.Clusters))
	}

	if sa.Clusters[0].Model.ID != 1 {
		t.Errorf("incorrect clusters ID: expected %d, got %d\n", 1, sa.Clusters[0].Model.ID)
	}

	// make sure data is correct
	expSA := &models.ServiceAccount{
		ProjectID:             1,
		Kind:                  "connector",
		AuthMechanism:         models.X509,
		ClientCertificateData: []byte("-----BEGIN"),
		ClientKeyData:         []byte("-----BEGIN"),
		Clusters: []models.Cluster{
			models.Cluster{
				ServiceAccountID:         1,
				Name:                     "cluster-test",
				Server:                   "https://localhost",
				CertificateAuthorityData: []byte("-----BEGIN"),
			},
		},
	}

	copySA := sa

	// reset fields for reflect.DeepEqual
	copySA.Model = orm.Model{}
	copySA.ClientCertificateData, _ = repository.Decrypt(copySA.ClientCertificateData, tester.key)
	copySA.ClientKeyData, _ = repository.Decrypt(copySA.ClientKeyData, tester.key)
	copySA.Clusters[0].Model = orm.Model{}
	copySA.Clusters[0].CertificateAuthorityData, _ = repository.Decrypt(
		copySA.Clusters[0].CertificateAuthorityData,
		tester.key,
	)

	if diff := deep.Equal(copySA, expSA); diff != nil {
		t.Errorf("incorrect service account")
		t.Error(diff)
	}
}

func TestListServiceAccountsByProjectID(t *testing.T) {
	tester := &tester{
		dbFileName: "./porter_list_sas.db",
	}

	setupTestEnv(tester, t)
	initProject(tester, t)
	initServiceAccount(tester, t)
	defer cleanup(tester, t)

	sas, err := tester.repo.ServiceAccount.ListServiceAccountsByProjectID(
		tester.initProjects[0].Model.ID,
	)

	if err != nil {
		t.Fatalf("%v\n", err)
	}

	if len(sas) != 1 {
		t.Fatalf("length of sas incorrect: expected %d, got %d\n", 1, len(sas))
	}

	// make sure data is correct
	expSA := &models.ServiceAccount{
		ProjectID:             1,
		Kind:                  "connector",
		AuthMechanism:         models.X509,
		ClientCertificateData: []byte("-----BEGIN"),
		ClientKeyData:         []byte("-----BEGIN"),
		Clusters: []models.Cluster{
			models.Cluster{
				ServiceAccountID:         1,
				Name:                     "cluster-test",
				Server:                   "https://localhost",
				CertificateAuthorityData: []byte("-----BEGIN"),
			},
		},
	}

	copySA := sas[0]

	// reset fields for reflect.DeepEqual
	copySA.Model = orm.Model{}
	copySA.ClientCertificateData, _ = repository.Decrypt(copySA.ClientCertificateData, tester.key)
	copySA.ClientKeyData, _ = repository.Decrypt(copySA.ClientKeyData, tester.key)
	copySA.Clusters[0].Model = orm.Model{}
	copySA.Clusters[0].CertificateAuthorityData, _ = repository.Decrypt(
		copySA.Clusters[0].CertificateAuthorityData,
		tester.key,
	)

	if diff := deep.Equal(copySA, expSA); diff != nil {
		t.Errorf("incorrect service account")
		t.Error(diff)
	}
}
