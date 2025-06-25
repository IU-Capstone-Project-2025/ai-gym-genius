import 'package:get_it/get_it.dart';

import '/core/domain/repositories/ex_infos_repository.dart';
import '/core/domain/repositories/training_repository.dart';
import '/core/data/repositories/workout_repository_impl.dart';
import '/core/data/repositories/ex_infos_repository_impl.dart';
import '/core/data/datasources/remote/remote_workout_datasource.dart';
import '/core/data/datasources/local/local_workout_datasource.dart';
import '/core/data/datasources/local/services/exercise_loader.dart';
import '/core/presentation/bloc/training_bloc.dart';

final GetIt getIt = GetIt.instance;

void setUpLocator() async {
  // Register services
  getIt.registerSingleton<ExerciseInfosLoader>(
    JsonExerciseInfosLoader(),
  );

  // Register Datasources
  // For local - dbProvider
  // For remote - apiProvider
  getIt.registerLazySingleton<LocalWorkoutDatasource>(
    () => SqfliteDatabase(),
  );
  getIt.registerLazySingleton<RemoteWorkoutDatasource>(
    () => APIWorkoutDatasource(),
  );

  // Register Repositories
  getIt.registerLazySingleton<TrainingRepository>(
    () => TrainingRepositoryImpl(),
  );
  getIt.registerLazySingleton<ExInfosRepository>(
    () => ExInfosRepositoryImpl(),
  );

  // Blocs
  getIt.registerFactory<TrainingBloc>(
    () => TrainingBloc(workoutRepository: getIt<TrainingRepository>()),
  );
}
