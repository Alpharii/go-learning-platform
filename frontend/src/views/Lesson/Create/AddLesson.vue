<template>
  <div class="flex justify-center items-start p-6">
    <Card class="w-full max-w-2xl">
      <CardHeader>
        <CardTitle>Tambah Lesson Baru</CardTitle>
      </CardHeader>

      <CardContent>
        <div v-if="errorLesson" class="text-red-500 mb-4 text-center">
          {{ errorLesson }}
        </div>

        <form @submit.prevent="handleLessonSubmit" class="space-y-6">
          <div class="space-y-2">
            <label for="lesson-title" class="block font-medium">Judul Lesson</label>
            <Input
              v-model="lessonTitle"
              id="lesson-title"
              placeholder="Masukkan judul lesson"
            />
          </div>

          <div class="space-y-2">
            <label for="lesson-content" class="block font-medium">
              Konten Lesson
            </label>
            <Textarea
              v-model="lessonContent"
              id="lesson-content"
              placeholder="Masukkan konten lesson"
              rows="5"
            />
          </div>

          <div class="space-y-2">
            <label for="lesson-order" class="block font-medium">Urutan</label>
            <Input
              type="number"
              v-model.number="lessonOrder"
              id="lesson-order"
              placeholder="Masukkan urutan"
            />
          </div>

          <div class="space-y-2">
            <label for="lesson-image" class="block font-medium">
              Gambar Lesson
            </label>
            <Input
              id="lesson-image"
              type="file"
              accept="image/*"
              @change="handleLessonImageUpload"
            />
          </div>

          <div v-if="lessonImagePreview" class="mt-4">
            <p class="font-medium mb-2">Pratinjau Gambar:</p>
            <img
              :src="lessonImagePreview"
              alt="Preview Gambar"
              class="w-full rounded-lg object-cover"
            />
          </div>

          <div class="flex justify-end">
            <Button
              type="submit"
              :disabled="isSubmittingLesson"
              class="flex items-center gap-2"
            >
              <Loader2
                v-if="isSubmittingLesson"
                class="animate-spin h-5 w-5"
              />
              <Upload v-else class="h-5 w-5" />
              <span>
                {{ isSubmittingLesson ? "Menyimpan..." : "Tambah Lesson" }}
              </span>
            </Button>
          </div>
        </form>
      </CardContent>
    </Card>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import { useRoute } from "vue-router";
import { createLesson } from "@/services/lessonService";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Loader2, Upload } from "lucide-vue-next";
import { useCourseStore } from "@/stores/course";

const route = useRoute();

const courseStore = useCourseStore()
const courseId = ref(courseStore.currentCourseId)

const lessonTitle = ref("");
const lessonContent = ref("");
// inisialisasi null supaya validasi >0 bisa jalan
const lessonOrder = ref<number | null>(null);
const lessonImageFile = ref<File | null>(null);
const lessonImagePreview = ref<string | null>(null);

const errorLesson = ref<string | null>(null);
const isSubmittingLesson = ref(false);

const handleLessonImageUpload = (e: Event) => {
  const files = (e.target as HTMLInputElement).files;
  if (files?.length) {
    lessonImageFile.value = files[0];
    lessonImagePreview.value = URL.createObjectURL(files[0]);
  }
};

const handleLessonSubmit = async () => {
  if (
    !lessonTitle.value ||
    !lessonContent.value ||
    !lessonOrder.value ||
    lessonOrder.value <= 0 ||
    !courseId.value
  ) {
    console.log(lessonTitle.value, lessonContent.value, lessonOrder.value, courseId.value)
    alert("Judul, Konten, Urutan, dan Course ID wajib diisi!");
    console.log('id', courseId.value)
    return;
  }

  isSubmittingLesson.value = true;
  const formData = new FormData();
  formData.append("title", lessonTitle.value);
  formData.append("content", lessonContent.value);
  formData.append("order", String(lessonOrder.value));
  formData.append("course_id", courseId.value);

  if (lessonImageFile.value) {
    formData.append("image", lessonImageFile.value);
  }

  try {
    await createLesson(formData);
    alert("Lesson berhasil ditambahkan!");
    // reset form
    lessonTitle.value = "";
    lessonContent.value = "";
    lessonOrder.value = null;
    lessonImageFile.value = null;
    lessonImagePreview.value = null;
  } catch (err) {
    errorLesson.value = "Gagal menambahkan lesson.";
    console.error(err);
  } finally {
    isSubmittingLesson.value = false;
  }
};
</script>
