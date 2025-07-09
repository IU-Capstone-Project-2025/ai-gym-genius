import 'package:get_it/get_it.dart';
import 'package:gym_genius/core/data/datasources/local/services/workout_database_provider.dart';
import 'package:gym_genius/core/data/repositories/mock_workout_repository_impl.dart';
import 'package:sqflite/sqflite.dart';

import '/core/domain/repositories/ex_infos_repository.dart';
import 'core/domain/repositories/workout_repository.dart';
import '/core/data/repositories/workout_repository_impl.dart';
import '/core/data/repositories/ex_infos_repository_impl.dart';
import '/core/data/datasources/remote/remote_workout_datasource.dart';
import '/core/data/datasources/local/local_workout_datasource.dart';
import '/core/data/datasources/local/services/exercise_loader.dart';
import '/core/presentation/bloc/training_bloc.dart';

final GetIt getIt = GetIt.instance;

/// Indicates whether to use mock versions or not.
enum LaunchingType {
  development,
  production,
}

void setUpLocator(LaunchingType type) async {
  // Register services
  getIt.registerSingleton<ExerciseInfosLoader>(
    JsonExerciseInfosLoader(),
  );

  // Register Datasources
  // For local - dbProvider
  // For remote - apiProvider
  getIt.registerLazySingleton<LocalWorkoutDatasource>(
    () => SqfliteDatabase(WorkoutDatabaseProvider(JsonExerciseInfosLoader(), databaseFactory)),
  );
  getIt.registerLazySingleton<RemoteWorkoutDatasource>(
    () => APIWorkoutDatasource(),
  );

  // Register Repositories
  getIt.registerLazySingleton<WorkoutRepository>(
    () {
      switch (type) {
        case LaunchingType.development:
          return MockWorkoutRepositoryImpl();
        case LaunchingType.production:
          return WorkoutRepositoryImpl();
      }
    },
  );
  getIt.registerLazySingleton<ExInfosRepository>(
    () => ExInfosRepositoryImpl(),
  );

  // Blocs
  getIt.registerFactory<TrainingBloc>(
    () => TrainingBloc(workoutRepository: getIt<WorkoutRepository>()),
  );
}
